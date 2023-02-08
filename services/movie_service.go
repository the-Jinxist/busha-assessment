package services

import "github.com/the-Jinxist/busha-assessment/services/models"

type MovieService interface {
	GetMovies() (*models.MovieResponse, error)
	GetMovieCharacters(movieID string) (*models.CharactersResponse, error)
}
