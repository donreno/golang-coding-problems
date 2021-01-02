package arrays

import (
	"unicode"
)

// IsPalindromePermutation checks if a given string is a permutation of a palindrome
func IsPalindromePermutation(s string) bool {
	if len(s) == 0 {
		return false
	}

	sFreq := buildFreqMatrixWithoutWhitespaces(s)

	countUnevenOcurrences := 0

	for _, count := range sFreq {
		if countUnevenOcurrences > 1 {
			return false
		}

		if count%2 != 0 {
			countUnevenOcurrences++
		}
	}

	return true
}

func buildFreqMatrixWithoutWhitespaces(s string) map[rune]uint {
	sFreq := make(map[rune]uint)

	for _, c := range s {
		if c != rune(' ') {
			c = unicode.ToLower(c)
			sFreq[c] = sFreq[c] + 1
		}
	}

	return sFreq
}
