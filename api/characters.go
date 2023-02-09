package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/the-Jinxist/busha-assessment/services/models"
	"github.com/the-Jinxist/busha-assessment/util"
)

type CharacterAPIResponse struct {
	Characters []*models.CharactersResponse `json:"characters"`
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusBadRequest))
		return
	}

	if len(request.Order) > 0 && len(request.Sort) == 0 {
		err = fmt.Errorf("characters cannot be ordered without a sort key")
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	redisKey, err := json.Marshal(request)

	if err != nil {
		log.Printf("error while marshalling request: %s", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusBadRequest))
		return
	}

	response := CharacterAPIResponse{}
	err = s.redisClient.Get(ctx, string(redisKey)).Scan(&response)
	if err != nil {
		if err != redis.Nil {
			log.Printf("error while getting redis value: %s", err)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
			return
		}
	}

	if len(response.Characters) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   response,
		})

		return
	}

	//if not, make the call to the swapi service
	var finalCharacters []*models.CharactersResponse
	characters, err := s.MovieService.GetMovieCharacters((strconv.Itoa(int(request.MovieID))))
	if err != nil {
		log.Printf("error while getting movie characters: %s", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusBadRequest))
		return
	}

	//parse the sort/filter elements if any
	if len(request.Sort) > 0 {
		sortCharacters(request.Sort, request.Order, characters)
	}

	finalCharacters = characters

	//filter the list
	if len(request.FilterBy) > 0 {
		filteredCharacters := filterCharacters(request.FilterBy, characters)
		finalCharacters = filteredCharacters
	}

	//save to redis using a combination of the endpoint and the params
	response.Characters = finalCharacters
	s.redisClient.Set(ctx, string(redisKey), response, time.Hour)

	//send to user
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   response,
	})
}

func sortCharacters(sortKey string, orderKey string, characters []*models.CharactersResponse) {

	if sortKey == util.HEIGHT {
		sort.Slice(characters, func(i, j int) bool {
			height1, err := strconv.Atoi(characters[i].Height)
			if err != nil {
				height1 = 0
			}

			height2, err := strconv.Atoi(characters[j].Height)
			if err != nil {
				height2 = 0
			}

			if orderKey == util.ASC {
				return height1 < height2
			}

			if orderKey == util.DESC {
				return height1 > height2
			}

			return height1 > height2

		})

		return
	}

	if sortKey == util.GENDER {
		sort.Slice(characters, func(i, j int) bool {
			character1 := characters[i]
			character2 := characters[j]

			if orderKey == util.ASC {
				return characterSortingWhenAsc(character1, character2)
			}

			if orderKey == util.DESC {
				return characterSortingWhenDsc(character1, character2)
			}

			return characterSortingWhenAsc(character1, character2)

		})

		return
	}

	if sortKey == util.NAME {
		sort.Slice(characters, func(i, j int) bool {
			character1 := characters[i]
			character2 := characters[j]

			if orderKey == util.ASC {
				return character1.Name < character2.Name
			}

			if orderKey == util.DESC {
				return character1.Name > character2.Name
			}

			return character1.Name < character2.Name

		})

		return
	}

}

func characterSortingWhenAsc(character1 *models.CharactersResponse, character2 *models.CharactersResponse) bool {
	if character1.Gender == "male" && character2.Gender == "female" {
		return true
	}

	if character1.Gender == "female" && character2.Gender == "male" {
		return false
	}

	if character1.Gender != "n/a" && character2.Gender == "n/a" {
		return true
	}

	if character1.Gender == character2.Gender {
		return true
	}

	return true
}

func characterSortingWhenDsc(character1 *models.CharactersResponse, character2 *models.CharactersResponse) bool {
	if character1.Gender == "male" && character2.Gender == "female" {
		return false
	}

	if character1.Gender == "female" && character2.Gender == "male" {
		return true
	}

	if character1.Gender != "n/a" && character2.Gender == "n/a" {
		return false
	}

	if character1.Gender == character2.Gender {
		return false
	}

	return false
}

func filterCharacters(filterKey string, characters []*models.CharactersResponse) []*models.CharactersResponse {

	filteredList := make([]*models.CharactersResponse, 0, 7)

	for index := range characters {
		if characters[index].Gender == filterKey {
			filteredList = append(filteredList, characters[index])
		}
	}

	return filteredList

}
