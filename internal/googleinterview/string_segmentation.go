package googleinterview

// IsStringSegmentation can the word be segmented in dictionary words
func IsStringSegmentation(word string, dictionary map[string]bool) bool {
	return isStringSegmentation([]rune(word), dictionary)
}

// TODO: continue checking string problems

func isStringSegmentation(word []rune, dictionary map[string]bool) bool {
	for i := range word {
		first := word[0:i]
		if dictionary[string(first)] {
			second := word[i:]
			if len(second) == 0 || dictionary[string(second)] || isStringSegmentation(second, dictionary) {
				return true
			}
		}
	}

	return false
}
