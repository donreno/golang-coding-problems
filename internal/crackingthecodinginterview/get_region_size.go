package crackingthecodinginterview

// GetMaxRegionSize get max region size among regions on the matrix
func GetMaxRegionSize(matrix [][]int) int {
	maxRegion := 0

	for r := range matrix {
		for c := range matrix[r] {
			regionSize := getRegionSize(matrix, r, c)
			maxRegion = max(maxRegion, regionSize)
		}
	}

	return maxRegion
}

func getRegionSize(matrix [][]int, row, col int) int {
	if !isInBounds(matrix, row, col) {
		return 0
	}

	if matrix[row][col] == 0 {
		return 0
	}

	matrix[row][col] = 0
	size := 1
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			size += getRegionSize(matrix, r, c)
		}
	}

	return size
}

func isInBounds(matrix [][]int, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[row])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
