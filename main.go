package main

import (
	"database/sql"
	"log"

	"github.com/the-Jinxist/busha-assessment/api"
	"github.com/the-Jinxist/busha-assessment/database/cache"
	database "github.com/the-Jinxist/busha-assessment/database/sqlc"
	"github.com/the-Jinxist/busha-assessment/services"
	"github.com/the-Jinxist/busha-assessment/util"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
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

	runDBMigrations("github.com://the-Jinxist/busha-assessment/database/migration#main", config.DBSource)

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

	err = server.Start(":" + config.ServerAddress)
	if err != nil {
		log.Fatalf("start server: %s", err)
	}
}

func runDBMigrations(migrationURL string, dbSourceString string) {

	log.Printf("migrationURL: %s, dbSourceString: %s", migrationURL, dbSourceString)

	migration, err := migrate.New(dbSourceString, migrationURL)
	if err != nil {
		log.Fatalf("cannot create new migrate instance: %s", err.Error())
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("cannot run up migrations: %s", err.Error())
	}

	log.Fatalf("db migrated successfully")
}
