package bits

// FlipToWin flip exactly one bit
func FlipToWin(in int) int {
	maxbitsCount, currentMaxCount, previousMaxCount := 1, 0, 0

	for in > 0 {
		if in&1 == 1 {
			currentMaxCount++
		} else if in&1 == 0 {
			if in&2 == 0 {
				previousMaxCount = 0
			} else {
				previousMaxCount = currentMaxCount
			}
			currentMaxCount = 0
		}

		if previousMaxCount+currentMaxCount > maxbitsCount {
			maxbitsCount = previousMaxCount + currentMaxCount
		}

		in = in >> 1
	}

	return maxbitsCount + 1
}
