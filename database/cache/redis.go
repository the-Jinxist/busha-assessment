package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/the-Jinxist/busha-assessment/util"
)

func NewRedis(config util.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
