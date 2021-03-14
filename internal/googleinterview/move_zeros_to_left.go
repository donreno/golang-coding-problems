package googleinterview

// MoveZerosToLeft moves all zeroes to left
func MoveZerosToLeft(arr []int) []int {
	lastIndex, firstZeroIndex, lastZeroIndex := len(arr)-1, -1, -1

	for lastIndex >= 0 && firstZeroIndex != 0 {
		currentVal := arr[lastIndex]
		if firstZeroIndex == -1 && lastZeroIndex == -1 && currentVal == 0 {
			firstZeroIndex, lastZeroIndex = lastIndex, lastIndex
		} else if currentVal == 0 {
			firstZeroIndex--
		} else if lastZeroIndex != -1 && currentVal != 0 {
			swap(arr, lastIndex, lastZeroIndex)
			lastZeroIndex--
			firstZeroIndex--
		}

		lastIndex--
	}

	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
