package recursionndynamic

// RobotInAGrid returns true if robot was able to get out
func RobotInAGrid(grid [][]int) bool {
	r, c := 0, 0
	lastR, lastC := len(grid)-1, len(grid[0])-1

	// assuming matrix is symmetric
	for r < lastR && c < lastC {
		checkDeadEnds(grid, r+1, c)
		checkDeadEnds(grid, r, c+1)

		if isDeadEnd(grid, r, c) {
			return false
		}

		if r < lastR && grid[r+1][c] != 1 {
			r++
		} else if c < lastC && grid[r][c+1] != 1 {
			c++
		}
	}

	return true
}

func checkDeadEnds(m [][]int, r, c int) bool {
	rowLen := len(m)
	colLen := len(m[0])

	//If out of bounds
	if r >= rowLen || c >= colLen {
		return false
	}

	//If I already stepped in this block
	if m[r][c] == 1 {
		return true
	}

	if isDeadEnd(m, r, c) {
		m[r][c] = 1
		return true
	}

	return checkDeadEnds(m, r+1, c) && checkDeadEnds(m, r, c+1)
}

//TODO: Check is dead end
func isDeadEnd(m [][]int, r, c int) bool {
	rowLen := len(m)
	colLen := len(m[0])

	if r == rowLen-1 || c == colLen-1 {
		return false
	}

	return (m[r+1][c] == 1 && m[r][c+1] == 1)
}
