package arrays

import (
	"strconv"
	"strings"
)

// CompressString returns a compressed version of the string if it's smaller than the original one
func CompressString(s string) string {
	sLen := len(s)
	if sLen <= 2 {
		return s
	}

	compressedBuffer := new(strings.Builder)
	compressed := ""
	sRunes := []rune(s)
	countCurrentRune := 1
	currentChar := sRunes[0]

	for i := 1; i < sLen; i++ {
		if i == sLen-1 {
			countCurrentRune++
		}

		if sRunes[i] != currentChar || i == sLen-1 {
			compressedBuffer.WriteRune(currentChar)
			compressedBuffer.WriteString(strconv.Itoa(countCurrentRune))
			currentChar = sRunes[i]
			countCurrentRune = 1
		} else {
			countCurrentRune++
		}
	}

	compressed = compressedBuffer.String()

	if sLen < len(compressed) {
		return s
	}

	return compressed
}
