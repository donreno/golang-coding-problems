package sorting

import "strings"

// SparseSearch search sorted string with empty strings in between
func SparseSearch(strs []string, elem string) int {
	low := 0
	high := len(strs) - 1

	for low <= high {
		mid := int((low + high) / 2)

		left := mid - 1
		right := mid + 1
		for strs[mid] == "" {

			if left < low && right > high {
				return -1
			} else if right <= high && strs[right] != "" {
				mid = right
			} else if left >= low && strs[left] != "" {
				mid = left
			}

			left--
			right++
		}

		middle := strs[mid]

		if middle == elem {
			return mid
		} else if strings.Compare(middle, elem) < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
