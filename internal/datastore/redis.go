package datastore

import (
	"github.com/redis/go-redis/v9"
	"im-GIN/config"
)

var Cache *redis.Client

func InitRedis(redisCfg *config.Redis) {
	addr := redisCfg.Host + ":" + redisCfg.Port
	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
}
