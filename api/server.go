package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	database "github.com/the-Jinxist/busha-assessment/database/sqlc"
	"github.com/the-Jinxist/busha-assessment/services"
	"github.com/the-Jinxist/busha-assessment/util"
)

// This struct [Server] will serve all our HTTP requests for our banking services
type Server struct {
	store        database.Store
	router       *gin.Engine
	config       util.Config
	MovieService services.MovieService
}

func NewServer(config util.Config, store database.Store, movieService services.MovieService) (*Server, error) {

	server := &Server{
		store:        store,
		config:       config,
		MovieService: movieService,
	}

	server.serveRouter()

	//Here, we register the custom validator, we created for validating currency here
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validSortType", validSortType)
		v.RegisterValidation("validOrder", validOrder)
		v.RegisterValidation("validGenderFilter", validGenderFilter)
	}

	return server, nil
}

func (server *Server) serveRouter() {

	router := gin.Default()

	router.GET("/movies", server.getMovies)
	router.POST("/comment", server.postComment)
	router.GET("/comments", server.getComments)
	router.GET("/characters", server.getCharacters)

	server.router = router
}

// Startes the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Startes the HTTP server on a specific address
func (server *Server) GetRedisClient() (*redis.Client, error) {
	service, ok := server.MovieService.(*services.SwapiService)
	if !ok {
		log.Println("error while casting MovieService to SwapiService")
		return nil, fmt.Errorf("error while casting MovieService to SwapiService")
	}

	log.Printf("SwapiServer: %v", service)
	return service.RedisClient, nil
}

func errorResponse(err error, status int) gin.H {
	return gin.H{
		"status": status,
		"error":  err,
	}
}
