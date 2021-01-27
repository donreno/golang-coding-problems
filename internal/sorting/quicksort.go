package sorting

// QuickSort quick sorts an array
func QuickSort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, left, right int) {
	index := partition(arr, left, right)

	if left < index-1 {
		quicksort(arr, left, index-1)
	}

	if index < right {
		quicksort(arr, index, right)
	}
}

func partition(arr []int, left, right int) int {
	pivotIndex := int((left + right) / 2)

	for left <= right {
		for arr[left] < arr[pivotIndex] {
			left++
		}

		for arr[right] > arr[pivotIndex] {
			right--
		}

		if left <= right {
			//Swaps elements
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return left
}
