package datastore

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"im-GIN/conf"
)

var Cache *redis.Client

func InitCache() {
	addr := conf.Cfg.Redis.Host + ":" + conf.Cfg.Redis.Port
	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Cfg.Redis.Password,
		DB:       conf.Cfg.Redis.DB,
	})
}

func GetAccessKey(uid uint64) string {
	return fmt.Sprintf("access_%d", uid)
}
