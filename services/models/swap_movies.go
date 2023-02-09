package models

import (
	"encoding/json"
	"time"
)

type SwapMovieList struct {
	Count       int          `json:"count"`
	Next        interface{}  `json:"next"`
	Previous    interface{}  `json:"previous"`
	SwapiMovies []SwapiMovie `json:"SwapiMovie"`
}

type SwapiMovie struct {
	Title        string    `json:"title"`
	EpisodeID    int       `json:"episode_id"`
	OpeningCrawl string    `json:"opening_crawl"`
	Director     string    `json:"director"`
	Producer     string    `json:"producer"`
	ReleaseDate  string    `json:"release_date"`
	Characters   []string  `json:"characters"`
	Planets      []string  `json:"planets"`
	Starships    []string  `json:"starships"`
	Vehicles     []string  `json:"vehicles"`
	Species      []string  `json:"species"`
	Created      time.Time `json:"created"`
	Edited       time.Time `json:"edited"`
	URL          string    `json:"url"`
}

func (i SwapiMovie) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func (i *SwapiMovie) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}
