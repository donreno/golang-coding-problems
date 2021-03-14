package googleinterview

// SumOfTwo sum of two func type
type SumOfTwo func([]int, int) bool

// SumOfTwoUnsorted with unsorted array
func SumOfTwoUnsorted(arr []int, target int) bool {
	complements := make(map[int]bool)

	for _, val := range arr {
		if complements[val] {
			return true
		}

		complements[target-val] = true
	}

	return false
}

// SumOfTwoSorted sum of two with sorted array
func SumOfTwoSorted(arr []int, target int) bool {
	startIndex, endIndex := 0, len(arr)-1

	for startIndex < endIndex {
		currentVal := arr[startIndex] + arr[endIndex]
		if currentVal < target {
			startIndex++
		} else if currentVal > target {
			endIndex--
		} else {
			return true
		}
	}

	return false
}
