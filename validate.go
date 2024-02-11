package go_nuban

const (
	nuban_length = 10
)

func Validate(nuban string) bool {
	if len(nuban) != nuban_length {
		return false // NUBAN code length must be 10
	}

	// Extract financial institution code and serial number
	fiCode := nuban[:6]
	serialNumber := nuban[6:9]
	checkDigit := int(nuban[9] - '0')

	// Calculate expected check digit
	expectedCheckDigit := calculateCheckDigit(fiCode + serialNumber)

	// Validate check digit
	return checkDigit == expectedCheckDigit
}
