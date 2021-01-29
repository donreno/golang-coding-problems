package sorting

// Listy representing a list
type Listy []int

func (l Listy) elementAt(i int) int {
	if i < len(l) {
		return l[i]
	}

	return -1
}

// SortedSearchNoSize finds an element on a list with no size
func SortedSearchNoSize(l Listy, elem int) int {
	//Find len

	var lastIndex int = 1
	for l.elementAt(lastIndex) != -1 && l.elementAt(lastIndex) < elem {
		lastIndex *= 2
	}

	//Binary Search
	return listyBinarySearch(l, elem, lastIndex/2, lastIndex)
}

func listyBinarySearch(l Listy, elem, low, high int) int {
	for low <= high {
		mid := int((low + high) / 2)
		middle := l.elementAt(mid)
		if middle > elem || middle == -1 {
			high = mid - 1
		} else if middle < elem {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
