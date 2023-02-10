package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/the-Jinxist/busha-assessment/services/models"
	"github.com/the-Jinxist/busha-assessment/util"
)

type Metadata struct {
	CharacterNumber       int    `json:"total_number_of_characters"`
	TotalHeightInCm       string `json:"total_height_in_cm"`
	TotalHeightInFeetInch string `json:"total_height_in_feet_inch"`
}

type CharacterAPIResponse struct {
	Characters        []*models.CharactersResponse `json:"characters"`
	CharacterMetadata *Metadata                    `json:"character_metadata"`
}

func (i CharacterAPIResponse) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func (i *CharacterAPIResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

type CharacterAPIRequest struct {
	Sort     string `form:"sort_by" binding:"validSortType"`
	Order    string `form:"order_by" binding:"validOrder"`
	FilterBy string `form:"filter_by" binding:"validGenderFilter"`
	MovieID  int64  `form:"movie_id" binding:"required,max=6"`
}

func (s *Server) getCharacters(ctx *gin.Context) {

	var request CharacterAPIRequest
	err := ctx.ShouldBindQuery(&request)

	if err != nil {
		log.Printf("error while binding query: %s", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(ReformatValidationError(err), http.StatusBadRequest))
		return
	}

	if len(request.Order) > 0 && len(request.Sort) == 0 {
		err = fmt.Errorf("characters cannot be ordered without a sort key")
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	response := CharacterAPIResponse{}

	var finalCharacters []*models.CharactersResponse
	characters, err := s.MovieService.GetMovieCharacters((strconv.Itoa(int(request.MovieID))))
	if err != nil {
		log.Printf("error while getting movie characters: %s", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusBadRequest))
		return
	}

	//parse the sort/filter elements if any
	if len(request.Sort) > 0 {
		SortCharacters(request.Sort, request.Order, characters)
	}

	finalCharacters = characters

	//filter the list
	if len(request.FilterBy) > 0 {
		filteredCharacters := FilterCharacters(request.FilterBy, characters)
		finalCharacters = filteredCharacters
	}
	metaData := &Metadata{}
	totalHeight := 0

	for index := range finalCharacters {
		character := finalCharacters[index]
		height, err := strconv.Atoi(character.Height)

		if err != nil {

			height = 0
			log.Printf("error while parsing height: %s", err)
		}

		totalHeight += height
	}

	metaData.CharacterNumber = len(finalCharacters)
	metaData.TotalHeightInCm = fmt.Sprintf("%dcm", totalHeight)

	feetInch, err := util.ConvertCMHeightToFtInch(totalHeight)
	if err != nil {
		metaData.TotalHeightInFeetInch = "0ft/0inch"
		log.Printf("error while ConvertCMHeightToFtInch: %s", err)
	}
	metaData.TotalHeightInFeetInch = feetInch

	response.Characters = finalCharacters
	response.CharacterMetadata = metaData

	//send to user
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   response,
	})
}
