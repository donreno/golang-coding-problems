package crackingthecodinginterview

// Mergesort merge sort implementation
func Mergesort(arr []int) {
	mergeSort(arr, make([]int, len(arr)), 0, len(arr)-1)
}

func mergeSort(arr, temp []int, start, end int) {
	if start >= end {
		return
	}

	middle := int((start + end) / 2)
	mergeSort(arr, temp, start, middle)
	mergeSort(arr, temp, middle+1, end)
	mergeHalves(arr, temp, start, end)
}

func mergeHalves(arr, temp []int, start, end int) {
	leftEnd := int((start + end) / 2)
	rightStart := leftEnd + 1

	left, right, index := start, rightStart, start

	for left <= leftEnd && right <= end {
		if arr[left] < arr[right] {
			temp[index] = arr[left]
			left++
		} else {
			temp[index] = arr[right]
			right++
		}

		index++
	}

	// if there are still left elements this will complete to fill out temp
	for left <= leftEnd {
		temp[index] = arr[left]
		index++
		left++
	}

	// In the other hand if there are still right elements this will complete to fill out temp
	for right <= end {
		temp[index] = arr[right]
		index++
		right++
	}

	// copy elements to array
	for i := start; i <= end; i++ {
		arr[i] = temp[i]
	}
}
