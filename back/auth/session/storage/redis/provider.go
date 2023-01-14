package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/go-redis/redis/v9"
)

var pder = &Provider{}

const (
	sessionSet          = "sessions"
	sessionHashPrefixID = "sess:"
)

type Provider struct {
	lock sync.Mutex
}

func getHashID(sid string) string {
	return fmt.Sprintf("%s%s", sessionHashPrefixID, sid)
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	ctx := context.Background()

	values := make(map[string]interface{}, 2)
	now := time.Now()
	values["timeAccessed"] = now

	couple := redis.Z{
		Score:  float64(now.Unix()),
		Member: sid,
	}
	_, err := config.Conf.RedisClient.ZAdd(ctx, sessionSet, couple).Result()
	if err != nil {
		log.Logger.Error("couldn't add session to set" + err.Error())
		return nil, err
	}

	config.Conf.RedisClient.HSet(ctx, getHashID(sid), values)

	return &SessionStore{sid: sid}, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	ctx := context.Background()

	if score := config.Conf.RedisClient.ZScore(ctx, sessionSet, sid); score != nil {
		return &SessionStore{sid: sid}, nil
	}

	return pder.SessionInit(sid)
}

func (pder *Provider) SessionDestroy(sid string) error {
	ctx := context.Background()

	config.Conf.RedisClient.ZRem(ctx, sessionSet, sid)
	config.Conf.RedisClient.Del(ctx, getHashID(sid))
	return nil
}

func (pder *Provider) SessionGC(maxLifeTime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	ctx := context.Background()
	var err error

	start := "0"
	end := fmt.Sprintf("%d", time.Now().Unix()-maxLifeTime)

	rangeBy := redis.ZRangeBy{
		Min: start,
		Max: end,
	}
	res, err := config.Conf.RedisClient.ZRangeByScore(ctx, sessionSet, &rangeBy).Result()
	if err != nil {
		log.Logger.Error(fmt.Sprintf("couldn't get timed out sessions from redis: %s", err))
		return
	}

	sessSIDS := make([]string, len(res))
	interfaces := make([]interface{}, len(res))
	for i, sid := range res {
		sessSIDS[i] = getHashID(sid)
		interfaces[i] = sid
	}
	config.Conf.RedisClient.Del(ctx, sessSIDS...)
	config.Conf.RedisClient.ZRem(ctx, sessionSet, interfaces...)
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	ctx := context.Background()
	var err error = nil

	z := redis.Z{
		Member: sid,
		Score:  float64(time.Now().Unix()),
	}
	_, err = config.Conf.RedisClient.ZAddXX(ctx, sessionSet, z).Result()
	return err
}

func Register() {
	log.Logger.Trace("Adding a redis session provider")
	session.Register("redis", pder)
}
