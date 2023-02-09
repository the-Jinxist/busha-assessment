package util

// Constants for all supported sort types
const (
	NAME   = "name"
	GENDER = "gender"
	HEIGHT = "height"
)

// function that checks if the sort type is used in the operation
func IsSortSupported(sortType string) bool {
	switch sortType {
	case NAME, GENDER, HEIGHT, "":
		return true
	}

	return false
}

// Constants for all supported orders
const (
	ASC  = "asc"
	DESC = "desc"
)

func IsOrderSupported(orderType string) bool {
	switch orderType {
	case ASC, DESC, "":
		return true
	}

	return false
}

// Constants for all supported gender filters
const (
	MALE   = "male"
	FEMALE = "female"
)

func IsGenderFilterSupported(genderType string) bool {
	switch genderType {
	case MALE, FEMALE, "":
		return true
	}

	return false
}
