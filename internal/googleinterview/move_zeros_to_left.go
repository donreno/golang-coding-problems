package googleinterview

// MoveZerosToLeft moves all zeroes to left
func MoveZerosToLeft(arr []int) []int {
	backwardIndex, lastZeroIndex := len(arr)-1, -1

	for backwardIndex >= 0 {
		currentVal := arr[backwardIndex]
		if lastZeroIndex == -1 && currentVal == 0 {
			lastZeroIndex = backwardIndex
		} else if lastZeroIndex != -1 && currentVal != 0 {
			swap(arr, backwardIndex, lastZeroIndex)
			lastZeroIndex--
		}

		backwardIndex--
	}

	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
