package go_nuban

import "fmt"

func GenerateSortCode(fiCode, stateCode, branchCode string) string {
	// Add leading zeros for DMBs and OFIs
	if len(fiCode) == 3 { // DMBs
		fiCode = fiCode
	} else if len(fiCode) == 5 { // OFIs
		fiCode = "9" + fiCode
	} else {
		return "" // Invalid financial institution code
	}

	// Calculate check digit
	checkDigit := calculateCheckDigit(fiCode + stateCode + branchCode)

	// Return complete sort code
	if len(fiCode) == 3 { // DMBs
		return fmt.Sprintf("%s-%s-%s-%d", fiCode, stateCode, branchCode, checkDigit)
	} else if len(fiCode) == 6 { // OFIs
		return fmt.Sprintf("%s-%s-%s-%d", fiCode, stateCode, branchCode, checkDigit)
	}
	return "" // Invalid financial institution code
}
