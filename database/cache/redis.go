package cache

import (
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/the-Jinxist/busha-assessment/util"
)

func NewRedis(config util.Config) *redis.Client {

	url, err := redis.ParseURL(config.RedisAddress)
	if err != nil {
		log.Fatalf("error occurred while creating redis instance: %s", err)
	}

	redisPassword := url.Password
	redisAddress := url.Addr

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB

	})

	return rdb
}
