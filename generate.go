package go_nuban

import (
	"fmt"
	"strconv"
	"strings"
)

func Generate(inst_code int, serial_number int) string {
	// Convert institution code and serial number to strings
	institutionCodeStr := fmt.Sprintf("%03d", inst_code)
	serialNumberStr := fmt.Sprintf("%09d", serial_number)

	// Concatenate institution code and serial number
	nubanWithoutCheckDigit := institutionCodeStr + serialNumberStr

	// Calculate check digit
	var total int
	for i, digitStr := range strings.Split(nubanWithoutCheckDigit, "") {
		digit, _ := strconv.Atoi(digitStr)
		multiplier := 3
		if (i+1)%3 == 0 {
			multiplier = 7
		}
		total += digit * multiplier
	}
	checkDigit := 10 - (total % 10)
	if checkDigit == 10 {
		checkDigit = 0
	}

	// Construct NUBAN account number
	nuban := nubanWithoutCheckDigit + strconv.Itoa(checkDigit)
	return nuban
}
