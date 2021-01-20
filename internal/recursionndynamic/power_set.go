package recursionndynamic

// GetAllSubsets gets all subsets from a given set
func GetAllSubsets(set []int) [][]int {
	return getAllSubsets(set, 0)
}

func getAllSubsets(set []int, index int) [][]int {
	var allSubsets [][]int

	if len(set) == index {
		allSubsets = [][]int{[]int{}}
	} else {
		allSubsets = getAllSubsets(set, index+1)
		item := set[index]

		moresubsets := make([][]int, 0)

		for _, subset := range allSubsets {
			newSubset := make([]int, 0)
			newSubset = append(newSubset, item)
			newSubset = append(newSubset, subset...)
			moresubsets = append(moresubsets, newSubset)
		}
		allSubsets = append(allSubsets, moresubsets...)
	}

	return allSubsets
}
