package go_nuban

import (
	"fmt"
)

func Generate(fiCode, serialNumber string) string {
	// Add leading zeros for DMBs and leading '9' for OFIs
	if len(fiCode) == 3 { // DMBs
		fiCode = "000" + fiCode
	} else if len(fiCode) == 5 { // OFIs
		fiCode = "9" + fiCode
	} else {
		return "" // Invalid financial institution code
	}

	// Calculate check digit
	checkDigit := calculateCheckDigit(fiCode + serialNumber)

	// Return complete NUBAN code
	return serialNumber + fmt.Sprintf("%d", checkDigit)
}
