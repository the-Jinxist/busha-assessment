package models

type CharactersResponse struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	URL       string `json:"url"`
}
