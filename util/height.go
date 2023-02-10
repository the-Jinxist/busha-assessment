package util

import (
	"fmt"
)

func ConvertCMHeightToFtInch(heightInCm int) (string, error) {

	feetInches := ""

	rawInch := float64(heightInCm) / 2.54
	numberOfFeet := rawInch / 12

	feetInches += fmt.Sprintf("%dft/", int(numberOfFeet))

	remainingInch := rawInch - (numberOfFeet * 12)

	feetInches += fmt.Sprintf("%finch", remainingInch)

	return feetInches, nil

}
