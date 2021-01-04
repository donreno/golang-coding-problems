package arrays

// ZeroMatrix zeroes rows and columns of a matrix if they have a zero on them
func ZeroMatrix(m [][]int) {
	if len(m) == 0 {
		return
	}

	rowsWithZero := make(map[int]bool)
	colsWithZero := make(map[int]bool)

	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[r]); c++ {
			if m[r][c] == 0 {
				rowsWithZero[r] = true
				colsWithZero[c] = true
			}
		}
	}

	for r := range rowsWithZero {
		blankRow(m, r)
	}

	for c := range colsWithZero {
		blankCol(m, c)
	}
}

func blankRow(m [][]int, rowIndex int) {
	for c := 0; c < len(m[rowIndex]); c++ {
		m[rowIndex][c] = 0
	}
}

func blankCol(m [][]int, colIndex int) {
	for r := 0; r < len(m); r++ {
		m[r][colIndex] = 0
	}
}
