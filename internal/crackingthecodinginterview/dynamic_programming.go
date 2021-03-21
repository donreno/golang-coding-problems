package crackingthecodinginterview

const (
	// BLOCKEDPATH represents a Blocked path
	BLOCKEDPATH = -1
)

// FindPathsDynamic finds an exit from
func FindPathsDynamic(grid [][]int) int {
	lastCol := len(grid[0]) - 1
	lastRow := len(grid) - 1
	// Exit is always on lower right corner
	upperFromExit := grid[lastRow-1][lastCol]
	if upperFromExit != BLOCKEDPATH {
		grid[lastRow-1][lastCol] = 1
	}

	for row := lastRow - 2; row >= 0; row-- {
		if grid[row][lastCol] == BLOCKEDPATH {
			continue
		} else if grid[row+1][lastCol] == BLOCKEDPATH {
			grid[row][lastCol] = 0
		} else {
			grid[row][lastCol] = grid[row+1][lastCol]
		}
	}

	leftFromExit := grid[len(grid)-1][len(grid[0])-2]
	if leftFromExit != BLOCKEDPATH {
		grid[len(grid)-1][len(grid[0])-2] = 1
	}

	for col := lastCol - 2; col >= 0; col-- {
		if grid[lastRow][col] == BLOCKEDPATH {
			continue
		} else if grid[lastRow][col+1] == BLOCKEDPATH {
			grid[lastRow][col] = 0
		} else {
			grid[lastRow][col] = grid[lastRow][col+1]
		}
	}

	for row := lastRow - 1; row >= 0; row-- {
		for col := lastCol - 1; col >= 0; col-- {
			if grid[row][col] == BLOCKEDPATH {
				continue
			}

			sum := 0

			// right ways
			if grid[row][col+1] > 0 {
				sum += grid[row][col+1]
			}

			if grid[row+1][col] > 0 {
				sum += grid[row+1][col]
			}

			grid[row][col] = sum
		}
	}

	return grid[0][0]
}
