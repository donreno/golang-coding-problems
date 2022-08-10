package crackingthecodinginterview

// BinarySearchRecursive recursive binary search implementation
func BinarySearchRecursive(arr []int, target int) bool {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target, left, right int) bool {
	if left > right {
		return false
	}

	mid := left + int((right-left)/2)
	current := arr[mid]

	if current == target {
		return true
	} else if current > target {
		return binarySearchRecursive(arr, target, left, mid-1)
	} else {
		return binarySearchRecursive(arr, target, mid+1, right)
	}
}

// BinarySearchNonRecursive non recursive binary search
func BinarySearchNonRecursive(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + int((right-left)/2)
		current := arr[mid]

		if current == target {
			return true
		} else if current > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
