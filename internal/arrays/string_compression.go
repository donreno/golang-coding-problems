package arrays

import "strconv"

// CompressString returns a compressed version of the string if it's smaller than the original one
func CompressString(s string) string {
	sLen := len(s)
	if sLen <= 2 {
		return s
	}

	compressed := ""
	sRunes := []rune(s)
	countCurrentRune := 1
	currentChar := sRunes[0]

	for i := 1; i < sLen; i++ {
		if i == sLen-1 {
			countCurrentRune++
		}

		if sRunes[i] != currentChar || i == sLen-1 {
			lastRuneCompress := string(currentChar) + strconv.Itoa(countCurrentRune)
			compressed = compressed + lastRuneCompress
			currentChar = sRunes[i]
			countCurrentRune = 1
		} else {
			countCurrentRune++
		}
	}

	if sLen < len(compressed) {
		return s
	}

	return compressed
}
