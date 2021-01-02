package arrays

// IsOneAway checks if two strings are zero or one change away
func IsOneAway(a, b string) bool {
	lenDiff := len(a) - len(b)

	// checks if they differ in more than one char (quick check)
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}

	aFreqs := makeLettersFreqMatrix(a)
	bFreqs := makeLettersFreqMatrix(b)

	countChangesAway := 0

	for index, count := range aFreqs {
		if countChangesAway > 1 {
			return false
		}

		diff := count - bFreqs[index]

		if diff < 0 {
			diff = diff * -1
		}

		countChangesAway = countChangesAway + diff
	}

	return countChangesAway <= 1
}

// Assuming it contains just lowercase letters
func makeLettersFreqMatrix(s string) []int {
	freqLetters := make([]int, 32)

	for _, char := range s {
		index := int(char - 'a')

		freqLetters[index] = freqLetters[index] + 1
	}

	return freqLetters
}
