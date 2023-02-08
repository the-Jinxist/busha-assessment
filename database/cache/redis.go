package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/the-Jinxist/busha-assessment/util"
)

func NewRedis(config util.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
