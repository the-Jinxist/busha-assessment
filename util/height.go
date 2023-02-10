package util

import (
	"fmt"
)

func ConvertCMHeightToFtInch(heightInCm int) (string, error) {

	feetInches := ""

	firstInch := float32(heightInCm) / 2.54

	feet := firstInch / 12
	feetInches += fmt.Sprintf("%dft and ", int(feet))

	nextInch := int(feet) * 12
	remainingInch := firstInch - float32(nextInch)

	feetInches += fmt.Sprintf("%.2f inches", float32(remainingInch))

	return feetInches, nil

}
