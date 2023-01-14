package redis

import (
	"context"
	"fmt"

	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/log"
)

type SessionStore struct {
	sid string // unique session id
}

func (st *SessionStore) Set(key, value interface{}) error {
	ctx := context.Background()

	_, err := config.Conf.RedisClient.HSet(ctx, getHashID(st.sid), key, value).Result()
	if err != nil {
		return err
	}

	return pder.SessionUpdate(st.sid)
}

func (st *SessionStore) Get(key interface{}) interface{} {
	err := pder.SessionUpdate(st.sid)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("couldn't update session: %s", err))
		return nil
	}
	ctx := context.Background()

	k := fmt.Sprintf("%v", key)
	res, err := config.Conf.RedisClient.HGet(ctx, getHashID(st.sid), k).Result()

	if err != nil {
		log.Logger.Error(fmt.Sprintf("couldn't get value from session: %s", err))
		return nil
	}
	return res
}

func (st *SessionStore) Delete(key interface{}) error {
	var err error
	ctx := context.Background()

	k := fmt.Sprintf("%v", key)
	if _, err = config.Conf.RedisClient.HDel(ctx, getHashID(st.sid), k).Result(); err != nil {
		return err
	}

	return pder.SessionUpdate(st.sid)
}

func (st *SessionStore) SessionID() string {
	return st.sid
}
