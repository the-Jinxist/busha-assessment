package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/the-Jinxist/busha-assessment/services/models"
)

const (
	BASE_URL = "https://swapi.dev/api/"
)

type SwapiService struct {
	RedisClient *redis.Client
}

func (service *SwapiService) getBaseURL() string {
	return BASE_URL
}

func (service *SwapiService) GetMovies() ([]*models.MovieResponse, error) {

	url := fmt.Sprintf("%s%s", service.getBaseURL(), "films/")
	response := &models.SwapMovieList{}

	err := makeHTTPRequest(url, response)
	if err != nil {
		log.Printf("error while making http request: %s", err)
		return nil, err
	}

	results := make([]*models.MovieResponse, 0, 7)

	for index := range response.SwapiMovies {

		movie := response.SwapiMovies[index]

		finalResponse := &models.MovieResponse{}
		finalResponse.Director = movie.Director
		finalResponse.EpisodeID = movie.EpisodeID
		finalResponse.OpeningCrawl = movie.OpeningCrawl
		finalResponse.Producer = movie.Producer
		finalResponse.ReleaseDate = movie.ReleaseDate
		finalResponse.Title = movie.Title
		finalResponse.URL = movie.URL

		results = append(results, finalResponse)
	}

	return results, nil

}

func (service *SwapiService) GetMovieCharacters(movieID string) ([]*models.CharactersResponse, error) {

	results := make([]*models.CharactersResponse, 0, 7)

	film, err := service.GetFilmFromID(movieID)
	if err != nil {
		log.Printf("error while getting film from id: %s", err)
		return nil, err
	}

	for index := range film.Characters {
		url := film.Characters[index]
		character, err := service.GetCharacterFromURL(url)

		if err != nil {
			log.Printf("error while getting character from url: %v", err)
			return nil, err
		}

		results = append(results, character)
	}

	return results, nil
}

func (service *SwapiService) GetCharacterFromURL(url string) (*models.CharactersResponse, error) {
	response := &models.SwapiCharacters{}

	err := service.RedisClient.Get(context.Background(), url).Scan(response)
	if err != nil {
		if err != redis.Nil {
			log.Printf("error while getting redis value: %s", err)
			return nil, err
		}
	}

	character := &models.CharactersResponse{}

	if len(response.Name) > 0 {

		character.BirthYear = response.BirthYear
		character.Gender = response.Gender
		character.Height = response.Height
		character.Name = response.Name
		character.URL = response.URL

		return character, nil
	}

	err = makeHTTPRequest(url, response)
	if err != nil {
		return nil, err
	}

	_, err = service.RedisClient.Set(context.Background(), url, response, time.Hour).Result()
	if err != nil {
		return nil, err
	}

	character.BirthYear = response.BirthYear
	character.Gender = response.Gender
	character.Height = response.Height
	character.Name = response.Name
	character.URL = response.URL

	return character, nil
}

func (service *SwapiService) GetFilmFromID(movieID string) (*models.SwapiMovie, error) {
	url := fmt.Sprintf("%s%s%s", service.getBaseURL(), "films/", movieID)
	response := &models.SwapiMovie{}

	err := service.RedisClient.Get(context.Background(), url).Scan(response)
	if err != nil {
		if err != redis.Nil {
			log.Printf("error while getting redis value: %s", err)
			return nil, err
		}
	}

	if len(response.Characters) > 0 {
		return response, nil
	}

	err = makeHTTPRequest(url, response)
	if err != nil {
		return nil, err
	}

	_, err = service.RedisClient.Set(context.Background(), url, response, time.Hour).Result()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func makeHTTPRequest(url string, response interface{}) error {
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("error while setting request header: %s", err)
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Printf("error client is doing: %s", err)

		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while ioutil is reading: %s", err)
		return err
	}

	buf := bytes.NewBuffer(body)
	err = json.NewDecoder(buf).Decode(response)
	if err != nil {
		log.Printf("error while decoding json: %s", err)
		return err
	}
	return nil
}
