package sorting

// BinarySearch search for an element on a sorted array and returns index
func BinarySearch(arr []int, elem int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := int((low + high) / 2)

		if arr[mid] < elem {
			low++
		} else if arr[mid] > elem {
			high--
		} else {
			return mid
		}
	}

	return -1
}

// BinarySearchRecursive recursive binary search
func BinarySearchRecursive(arr []int, elem int) int {
	return binarySearchRecursive(arr, elem, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, elem, low, high int) int {
	if low > high {
		return -1
	}

	mid := int((low + high) / 2)

	if arr[mid] < elem {
		return binarySearchRecursive(arr, elem, mid+1, high)
	} else if arr[mid] > elem {
		return binarySearchRecursive(arr, elem, low, mid-1)
	}

	return mid

}
