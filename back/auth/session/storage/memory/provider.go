package memory

import (
	"container/list"
	"sync"
	"time"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/utils/log"
)

var pder = &Provider{
	list: list.New(),
}

type Provider struct {
	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // sessions are saved in memory
	list     *list.List               // gc
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)
	newSess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newSess)
	pder.sessions[sid] = element
	return newSess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxLifeTime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func Register() {
	log.Logger.Trace("Adding a memory session provider")
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}
