package api

import (
	"sort"
	"strconv"

	"github.com/the-Jinxist/busha-assessment/services/models"
	"github.com/the-Jinxist/busha-assessment/util"
)

func SortCharacters(sortKey string, orderKey string, characters []*models.CharactersResponse) {

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

	if character1.Gender == "female" {
		return true
	}

	if character1.Gender == "male" && character2.Gender == "n/a" {
		return true
	}

	return false
}

func characterSortingWhenDsc(character1 *models.CharactersResponse, character2 *models.CharactersResponse) bool {
	if character1.Gender == "n/a" {
		return true
	}

	if character1.Gender == "male" && character2.Gender == "female" {
		return true
	}

	return false
}

func FilterCharacters(filterKey string, characters []*models.CharactersResponse) []*models.CharactersResponse {

	filteredList := make([]*models.CharactersResponse, 0, 7)

	for index := range characters {
		if characters[index].Gender == filterKey {
			filteredList = append(filteredList, characters[index])
		}
	}

	return filteredList

}
