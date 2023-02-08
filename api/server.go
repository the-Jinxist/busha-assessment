package api

import (
	"github.com/gin-gonic/gin"
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
	redisClient  *redis.Client
}

func NewServer(config util.Config, store database.Store, movieService services.MovieService, redis *redis.Client) (*Server, error) {

	server := &Server{
		store:        store,
		config:       config,
		MovieService: movieService,
		redisClient:  redis,
	}

	server.serveRouter()
	return server, nil
}

func (server *Server) serveRouter() {

	router := gin.Default()

	router.GET("/movies", server.getMovies)

	server.router = router
}

// Startes the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error, status int) gin.H {
	return gin.H{
		"status": status,
		"error":  err.Error(),
	}
}
