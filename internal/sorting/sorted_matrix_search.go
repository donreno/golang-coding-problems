package sorting

import (
	"math"
)

// FindElementInSortedMatrix finds element's position in the matrix
func FindElementInSortedMatrix(m [][]int, elem int) (int, int) {

	return findElementInSortedMatrix(m, elem, 0, 0, len(m)-1, len(m[0])-1)
}

func findElementInSortedMatrix(m [][]int, elem, origRow, origCol, destRow, destCol int) (int, int) {
	if !inBounds(m, origRow, origCol) || !inBounds(m, destRow, destCol) {
		return -1, -1
	}

	if m[origRow][origCol] == elem {
		return origRow, origCol
	} else if !isBefore(origRow, origCol, destRow, destCol) {
		return -1, -1
	}

	startRow, startCol := origRow, origCol
	diagDist := int(math.Min(float64(destRow-origRow), float64(destCol-origCol)))
	endRow, endCol := origRow+diagDist, origCol+diagDist

	// finds first midpoint in diagonal that's bigger than elem
	for isBefore(startRow, startCol, endRow, endCol) {
		midRow, midCol := int((startRow+endRow)/2), int((startCol+endCol)/2)

		if elem > m[midRow][midCol] {
			startRow = midRow + 1
			startCol = midCol + 1
		} else {
			endRow = midRow - 1
			endCol = midCol - 1
		}
	}

	return partitionAndSearch(m, elem, origRow, origCol, destRow, destCol, startRow, startCol)
}

func partitionAndSearch(m [][]int, elem, origRow, origCol, destRow, destCol, pivotRow, pivotCol int) (int, int) {
	lowerLeftOrigRow, lowerLeftOrigCol := pivotRow, origCol
	lowerLeftDestRow, lowerLeftDestCol := destRow, pivotCol-1
	upperRightOrigRow, upperRightOrigCol := origRow, pivotCol
	upperRightDestRow, upperRightDestCol := pivotRow-1, destCol

	lowerLeftRow, lowerLeftCol := findElementInSortedMatrix(m, elem, lowerLeftOrigRow, lowerLeftOrigCol, lowerLeftDestRow, lowerLeftDestCol)

	if lowerLeftRow == -1 || lowerLeftCol == -1 {
		return findElementInSortedMatrix(m, elem, upperRightOrigRow, upperRightOrigCol, upperRightDestRow, upperRightDestCol)
	}

	return lowerLeftRow, lowerLeftCol
}

func inBounds(m [][]int, r, c int) bool {
	return r >= 0 && c >= 0 && r < len(m) && c < len(m[0])
}

func isBefore(origRow, origCol, destRow, destCol int) bool {
	return origRow <= destRow && origCol <= destCol
}
