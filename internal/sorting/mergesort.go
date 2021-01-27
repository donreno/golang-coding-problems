package sorting

// MergeSort mergesort implementation
func MergeSort(arr []int) {
	mergeSort(arr, make([]int, len(arr)), 0, len(arr)-1)
}

func mergeSort(arr, temp []int, left, right int) {
	if left < right {
		middle := int((left + right) / 2)

		mergeSort(arr, temp, left, middle)
		mergeSort(arr, temp, middle+1, right)
		merge(arr, temp, left, middle, right)
	}
}

func merge(arr, temp []int, left, middle, right int) {
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}

	helperLeft := left
	helperRight := middle + 1
	current := left

	for helperLeft <= middle && helperRight <= right {
		if temp[helperLeft] <= temp[helperRight] {
			arr[current] = temp[helperLeft]
			helperLeft++
		} else {
			arr[current] = temp[helperRight]
			helperRight++
		}

		current++
	}

	remaining := middle - helperLeft

	for i := 0; i <= remaining; i++ {
		arr[current+i] = temp[helperLeft+i]
	}
}
