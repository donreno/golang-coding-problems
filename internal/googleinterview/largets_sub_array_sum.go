package googleinterview

// FindLargestSubArraySum find the largest sum for a sub array
func FindLargestSubArraySum(arr []int) int {
	maxCurrent, maxGlobal := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		maxCurrent = max(arr[i], arr[i]+maxCurrent)

		maxGlobal = max(maxCurrent, maxGlobal)
	}

	return maxGlobal
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
