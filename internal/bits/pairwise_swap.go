package bits

const (
	oddBits  int16 = 0xAA
	evenBits int16 = 0x55
)

// PairwiseSwap swaps oddBits for even and viceversa
func PairwiseSwap(n int16) int16 {
	return (n & oddBits >> 1) | (n & evenBits << 1)
}
