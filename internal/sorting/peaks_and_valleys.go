package sorting

// PeaksAndValleys sorts array so it gets peaks and valleys
func PeaksAndValleys(arr []int) {
	for i := 1; i < len(arr); i += 2 {
		if getMaxValueIndex(arr, i-1, i) != i {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}
	}
}

func getMaxValueIndex(arr []int, first, second int) int {
	if arr[first] > arr[second] {
		return first
	}

	return second
}
