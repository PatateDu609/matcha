package memory

import "time"

type SessionStore struct {
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value stored inside
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	return pder.SessionUpdate(st.sid)
}

func (st *SessionStore) Get(key interface{}) interface{} {
	_ = pder.SessionUpdate(st.sid)

	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	return pder.SessionUpdate(st.sid)
}

func (st *SessionStore) SessionID() string {
	return st.sid
}
