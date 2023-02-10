package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/the-Jinxist/busha-assessment/services/models"
)

type MovieAPIResponse struct {
	Title         string `json:"title"`
	EpisodeID     int    `json:"episode_id"`
	OpeningCrawl  string `json:"opening_crawl"`
	Director      string `json:"director"`
	Producer      string `json:"producer"`
	ReleaseDate   string `json:"release_date"`
	URL           string `json:"url"`
	CommentNumber int    `json:"comment_number"`
}

type RedisStruct struct {
	Movies []*models.MovieResponse `json:"movies"`
}

func (i RedisStruct) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func (i *RedisStruct) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

func (s *Server) getMovies(ctx *gin.Context) {

	var rawResponse []*models.MovieResponse

	rawResponse, err := s.MovieService.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(ReformatValidationError(err), http.StatusInternalServerError))
		return
	}

	apiResponse, err := s.GetFullAPIResponse(ctx, rawResponse)
	if err != nil {
		log.Printf("error while getting full api response: %s", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	if err != nil {
		log.Printf("error while saving redis value: %s", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   apiResponse,
	})

}

func (s *Server) GetFullAPIResponse(ctx *gin.Context, rawResponse []*models.MovieResponse) ([]*MovieAPIResponse, error) {
	result := make([]*MovieAPIResponse, 0, 7)
	for index := range rawResponse {
		movie := rawResponse[index]
		commentNumber, err := s.store.GetCommentNumber(ctx, strconv.Itoa(movie.EpisodeID))

		if err != nil {
			return nil, err
		}

		response := &MovieAPIResponse{}
		response.Director = movie.Director
		response.Title = movie.Title
		response.EpisodeID = movie.EpisodeID
		response.OpeningCrawl = movie.OpeningCrawl
		response.Producer = movie.Producer
		response.ReleaseDate = movie.ReleaseDate
		response.URL = movie.URL
		response.CommentNumber = int(commentNumber)

		result = append(result, response)

	}

	return result, nil

}
