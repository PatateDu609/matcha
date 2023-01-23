package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
)

var (
	GlobalSessions *Manager

	errProviderNotFound = errors.New("no provider with the given name")
)

type Manager struct {
	cookieName  string // Private cookie name
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("%s (%s)", errProviderNotFound, provideName)
	}
	return &Manager{
		provider:    provider,
		cookieName:  cookieName,
		maxlifetime: maxlifetime,
	}, nil
}

// CheckSessionCookie returns true if the session cookie is set, and false otherwise
func (manager *Manager) CheckSessionCookie(r *http.Request) bool {
	cookie, err := r.Cookie(manager.cookieName)
	return err == nil && cookie.Value != ""
}

func (manager *Manager) sessionID() string {
	uid, err := uuid.NewRandom()

	if err != nil {
		b := make([]byte, 36)
		if _, err := io.ReadFull(rand.Reader, b); err != nil {
			return ""
		}
		return base64.URLEncoding.EncodeToString(b)
	}
	return uid.String()
}

// SessionStart retrieves the session cookie from the http.Request provided, if
// the cookie is not present or if its value is empty, a new session is created
// and is sent to the user by the mean of the http.ResponseWriter object. If no
// http.ResponseWriter is provided, no new session will be created and nil will
// be returned.
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		if err != nil {
			log.Logger.Errorf("couldn't get cookie: %s", err)
			log.Logger.Warn("trying to init new session")
			err = nil
		}
		if w == nil {
			session = nil
			return
		}

		sid := manager.sessionID()
		session, err = manager.provider.SessionInit(sid)
		if err != nil {
			log.Logger.Errorf("couldn't init new session: %s", err)
		}
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxlifetime),
		}

		http.SetCookie(w, &cookie)
	} else {
		var sid string
		sid, err = url.QueryUnescape(cookie.Value)
		if err != nil {
			return
		}
		session, err = manager.provider.SessionRead(sid)
	}
	return
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		if err != nil {
			return err
		}
		return nil
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()

	if err = manager.provider.SessionDestroy(cookie.Value); err != nil {
		return err
	}
	expiration := time.Now()
	cookie = &http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiration,
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
	return err
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}
