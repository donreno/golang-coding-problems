package arrays

import (
	"unicode"
)

// IsPalindromePermutationBitwise checks if a given string is a permutation of a palindrome (this time a bitwise implementation)
func IsPalindromePermutationBitwise(s string) bool {
	if len(s) == 0 {
		return false
	}

	bitVector := buildBitVector(s)

	return bitVector == 0 || hasExactlyOneBitOn(bitVector)
}

func buildBitVector(s string) int {
	bitVector := 0

	for _, c := range s {
		if c != rune(' ') {
			c = unicode.ToLower(c)
			index := int(c - 'a')

			bitVector = toggleBit(bitVector, index)
		}
	}

	return bitVector
}

// Toggles on/off bit
func toggleBit(bitVector, index int) int {
	if index < 0 {
		return bitVector
	}

	mask := 1 << index

	if bitVector&mask == 0 {
		bitVector = bitVector | mask
	} else {
		bitVector = bitVector &^ mask
	}

	return bitVector
}

// Checks if number has just one bit on
func hasExactlyOneBitOn(bitVector int) bool {
	return bitVector&(bitVector-1) == 0
}
