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
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
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

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxlifetime),
		}

		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()

	_ = manager.provider.SessionDestroy(cookie.Value)
	expiration := time.Now()
	cookie = &http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiration,
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}
