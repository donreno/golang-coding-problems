package googleinterview

// IsValidNumber checks if a string is a valid number
func IsValidNumber(num string) bool {
	isDecimal := false
	numRunes := []rune(num)
	digit := numRunes[0]

	if !isNumber(digit) {
		return false
	}

	for i := 1; i < len(numRunes); i++ {
		digit = numRunes[i]

		if isPeriod(digit) && isDecimal {
			return false
		} else if isPeriod(digit) {
			isDecimal = true
		} else if !isNumber(digit) {
			return false
		}
	}

	return isNumber(digit)
}

func isNumber(digit rune) bool {
	return digit >= '0' && digit <= '9'
}

func isPeriod(digit rune) bool {
	return digit == '.'
}
