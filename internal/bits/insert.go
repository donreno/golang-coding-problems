package bits

// InsertMIntoNAtItoJ inserts M into N at expected positions
func InsertMIntoNAtItoJ(m, n, i, j int) int {
	allOnes := ^0

	left := allOnes << (j + 1)
	right := (1 << i) - 1
	// Mask will clean all bytes between i and j
	mask := left | right

	nClear := n & mask
	mShifted := m << i

	return nClear | mShifted
}
