package redis

import (
	"fmt"
	"go-web-staging/setting"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	rdb.Close()
}
