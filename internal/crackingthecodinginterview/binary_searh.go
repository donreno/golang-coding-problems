package crackingthecodinginterview

// BinarySearchRecursive recursive binary search implementation
func BinarySearchRecursive(arr []int, target int) bool {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target, left, right int) bool {
	if left > right {
		return false
	}

	mid := int((left + right) / 2)

	if arr[mid] == target {
		return true
	} else if target < arr[mid] {
		return binarySearchRecursive(arr, target, left, mid-1)
	} else {
		return binarySearchRecursive(arr, target, mid+1, right)
	}
}

// BinarySearchNonRecursive non recursive binary search
func BinarySearchNonRecursive(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := int((left + right) / 2)

		if arr[mid] == target {
			return true
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
