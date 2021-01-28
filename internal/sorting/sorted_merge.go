package sorting

import "errors"

// SortedMerge sorted merge
func SortedMerge(a, b []int) error {
	aLen := len(a)
	bLen := len(b)

	if bLen >= aLen {
		return errors.New("Invalid length arrays")
	}

	mergeArrs(a, b)

	sortedMergeSort(a, make([]int, aLen), 0, aLen-1)

	return nil
}

func mergeArrs(a, b []int) {
	aLen := len(a)
	bLen := len(b)
	aLastIndex := aLen - bLen - 1

	for i := 0; i < bLen; i++ {
		a[aLastIndex+1+i] = b[i]
	}
}

func sortedMergeSort(a, temp []int, left, right int) {
	if left < right {
		mid := int((left + right) / 2)

		sortedMergeSort(a, temp, left, mid)
		sortedMergeSort(a, temp, mid+1, right)
		mergeWithSortedMerge(a, temp, left, mid, right)
	}
}

func mergeWithSortedMerge(a, temp []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		temp[i] = a[i]
	}

	tempLeft, current := left, left
	tempRight := mid + 1

	for tempLeft <= mid && tempRight <= right {
		if temp[tempLeft] <= temp[tempRight] {
			a[current] = temp[tempLeft]
			tempLeft++
		} else {
			a[current] = temp[tempRight]
			tempRight++
		}

		current++
	}

	remainingElems := mid - tempLeft

	for i := 0; i <= remainingElems; i++ {
		a[current+i] = temp[tempLeft+i]
	}
}
