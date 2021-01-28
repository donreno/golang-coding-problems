package sorting

// SearchInRotatedArray Searches an element in a rotated sorted array
func SearchInRotatedArray(arr []int, target int) int {
	return searchInRotatedArray(arr, target, 0, len(arr)-1)
}

func searchInRotatedArray(arr []int, target, low, high int) int {
	if arr[low] == target {
		return low
	}

	if arr[high] == target {
		return high
	}

	if low >= high {
		return -1
	}

	mid := int((high + low) / 2)

	if arr[mid] == target {
		return mid
	}

	//  are we in a loop
	if arr[low] > arr[high] {
		left := searchInRotatedArray(arr, target, low, mid)

		if left != -1 {
			return left
		}

		return searchInRotatedArray(arr, target, mid+1, high)
	}

	if arr[low] < target && arr[mid] > target {
		return searchInRotatedArray(arr, target, low, mid)
	}

	return searchInRotatedArray(arr, target, mid+1, high)
}
