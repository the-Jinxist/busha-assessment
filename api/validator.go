package api

import (
	"github.com/the-Jinxist/busha-assessment/util"

	"github.com/go-playground/validator/v10"
)

// This methods creates a new function on the validator.Func interface that allows us to validate a field
// in our JSON validator
var validSortType validator.Func = func(fieldLevel validator.FieldLevel) bool {
	sortType, ok := fieldLevel.Field().Interface().(string)
	if ok {
		return util.IsSortSupported(sortType)
	}

	return false
}

var validOrder validator.Func = func(fieldLevel validator.FieldLevel) bool {
	orderType, ok := fieldLevel.Field().Interface().(string)
	if ok {
		return util.IsOrderSupported(orderType)
	}

	return false
}

var validGenderFilter validator.Func = func(fieldLevel validator.FieldLevel) bool {
	orderType, ok := fieldLevel.Field().Interface().(string)
	if ok {
		return util.IsGenderFilterSupported(orderType)
	}

	return false
}
