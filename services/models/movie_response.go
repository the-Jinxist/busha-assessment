package models

type MovieResponse struct {
	Title        string `json:"title"`
	EpisodeID    int    `json:"episode_id"`
	OpeningCrawl string `json:"opening_crawl"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
	ReleaseDate  string `json:"release_date"`
	URL          string `json:"url"`
}
