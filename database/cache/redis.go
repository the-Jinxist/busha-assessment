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

	rdb := redis.NewClient(&redis.Options{
		Addr:     url.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
