package sorting

import "errors"

// SortedMerge sorted merge
func SortedMerge(a, b []int) error {
	aLen := len(a)
	bLen := len(b)

	if bLen >= aLen {
		return errors.New("Invalid length arrays")
	}

	indexA := aLen - bLen - 1
	indexB := bLen - 1
	indexMerged := aLen - 1

	for indexB >= 0 {
		if indexA >= 0 && a[indexA] > b[indexB] {
			a[indexMerged] = a[indexA]
			indexA--
		} else {
			a[indexMerged] = b[indexB]
			indexB--
		}

		indexMerged--
	}

	return nil
}
