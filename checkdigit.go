package go_nuban

func calculateCheckDigit(nuban string) int {
	weights := []int{3, 7, 3, 3, 7, 3, 3, 7, 3, 3, 7, 3, 3, 7, 3}
	sum := 0
	for i, digit := range nuban {
		sum += int(digit-'0') * weights[i]
	}
	checkDigit := (10 - (sum % 10)) % 10
	return checkDigit
}
