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

	countAddChanges := 0
	countRemoveChanges := 0

	for index, count := range aFreqs {
		diff := count - bFreqs[index]

		if diff < 0 {
			diff = diff * -1
			countAddChanges = countAddChanges + diff
		} else {
			countRemoveChanges = countRemoveChanges + diff
		}
	}

	// One add edit
	if countAddChanges == 1 && countRemoveChanges == 0 {
		return true
	}

	// One remove edit
	if countRemoveChanges == 1 && countAddChanges == 0 {
		return true
	}

	// one exchange edit
	if countAddChanges == 1 && countRemoveChanges == 1 {
		return true
	}

	return countAddChanges == 0 && countRemoveChanges == 0
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
