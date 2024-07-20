package main

import (
	"fmt"
	"go-web-staging/internal/app"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis(cfg *app.RedisConfig) (err error) {

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

func CloseRedis() {
	rdb.Close()
}
