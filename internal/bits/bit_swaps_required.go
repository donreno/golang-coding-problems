package bits

// BitSwapsRequiered bit swaps required so a becomes b
func BitSwapsRequiered(a, b int) int {

	xorBits := a ^ b
	count := 0

	for ; xorBits != 0; xorBits >>= 1 {
		count += xorBits & 1
	}

	return count
}
