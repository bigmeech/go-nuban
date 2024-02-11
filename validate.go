package go_nuban

import (
	"strconv"
	"strings"
)

const (
	nuban_length = 16
)

func Validate(account_number string) bool {
	if len(account_number) != nuban_length {
		return false
	}

	// Extract digits from account number
	digits := strings.Split(account_number, "")

	// Convert each digit to integer
	var total int
	for i, digitStr := range digits[:len(digits)-1] {
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			return false
		}
		multiplier := 3
		if (i+1)%3 == 0 {
			multiplier = 7
		}
		total += digit * multiplier
	}

	// Calculate check digit
	checkDigit := 10 - (total % 10)
	if checkDigit == 10 {
		checkDigit = 0
	}

	// Validate check digit
	lastDigit, err := strconv.Atoi(digits[len(digits)-1])
	if err != nil {
		return false
	}
	return checkDigit == lastDigit
}
