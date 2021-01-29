package sorting

// Bitset type to handle bitsets (in byte chunks)
type Bitset []byte

// MakeBitset creates a new bitset of given size
func MakeBitset(size int) Bitset {
	bytearr := make([]byte, int(size/8)+1)
	return Bitset(bytearr)
}

// Get gets ith bit
func (b Bitset) Get(i int) bool {
	arrPos := i / 8
	bitPos := 1 << (i % 8)
	return (int(b[arrPos]) & bitPos) != 0
}

// Set sets ith bit
func (b Bitset) Set(i int) {
	arrPos := i / 8
	bitPos := 1 << (i % 8)
	b[arrPos] |= byte(bitPos)
}

// CheckDuplicates checks duplicates and returns duplicates on arrays
func CheckDuplicates(arr []int) map[int]bool {
	bs := MakeBitset(32000)
	dups := make(map[int]bool)

	for _, v := range arr {
		if bs.Get(v) {
			dups[v] = true
		} else {
			bs.Set(v)
		}
	}

	return dups
}
