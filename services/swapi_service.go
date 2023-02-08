package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/the-Jinxist/busha-assessment/services/models"
)

type SwapiService struct {
	// BASE_URL := ""
}

func (service *SwapiService) GetMovies() ([]*models.MovieResponse, error) {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://swapi.dev/api/films/", nil)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("error while setting request header: %s", err)
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Printf("error client is doing: %s", err)

		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	response := &models.SwapMovies{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while ioutil is reading: %s", err)
	}

	buf := bytes.NewBuffer(body)
	err = json.NewDecoder(buf).Decode(response)
	if err != nil {
		log.Printf("error while decoding json: %s", err)
	}

	results := make([]*models.MovieResponse, 0, 7)

	for index := range response.Results {

		movie := response.Results[index]

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

	// layout := "1977-05-25"
	// sort.Slice(results, func(i, j int) bool {

	// 	// values := strings.Split(results[i].ReleaseDate, "-")
	// 	movie1Time, err := time.Parse(layout, results[i].ReleaseDate)
	// 	if err != nil {
	// 		log.Printf("error while parsing: %s", err)
	// 		return false
	// 	}

	// 	movie2Time, err := time.Parse(layout, results[j].ReleaseDate)

	// 	if err != nil {
	// 		log.Printf("error while parsing: %s", err)
	// 		return false
	// 	}

	// 	return movie1Time.Before(movie2Time)
	// })

	return results, nil

}

func (service *SwapiService) GetMovieCharacters(movieID string) (*models.CharactersResponse, error) {
	//Get the single movie

	//Loop through character's list and get the character profile

	//Send the list
	return nil, nil
}
