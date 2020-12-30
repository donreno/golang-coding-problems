package golang_coding_problems

// HasUniqueChars checks if a string has only unique chars
func HasUniqueChars(s string) bool {
	sLen := len(s)

	if sLen == 0 {
		return false
	}

	uniqueChars := make(map[rune]bool)

	for _, c := range s {
		if uniqueChars[c] {
			return false
		}

		uniqueChars[c] = true
	}

	return true
}
