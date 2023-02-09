package main

import (
	"database/sql"
	"log"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/the-Jinxist/busha-assessment/api"
	"github.com/the-Jinxist/busha-assessment/database/cache"
	database "github.com/the-Jinxist/busha-assessment/database/sqlc"
	"github.com/the-Jinxist/busha-assessment/services"
	"github.com/the-Jinxist/busha-assessment/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("error while opening database: %s", err)
	}

	redisClient := cache.NewRedis(config)

	service := &services.SwapiService{
		RedisClient: redisClient,
	}

	store := database.NewStore(conn)

	runHTTPServer(config, store, service)
}

func runHTTPServer(config util.Config, store database.Store, movieService services.MovieService) {
	server, err := api.NewServer(config, store, movieService)

	if err != nil {
		log.Fatalf("cannot create server: %s", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("start server: %s", err)
	}
}
