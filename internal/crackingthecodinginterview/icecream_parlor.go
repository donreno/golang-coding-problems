package crackingthecodinginterview

import "sort"

// GetIcecreamChoices Ice cream parlol, gets 2 choices of icecream given certain amount of money
func GetIcecreamChoices(menu []int, money int) []int {
	complementsPosition := make(map[int]int)

	for i, cost := range menu {
		if complementPos, found := complementsPosition[cost]; found {
			return []int{complementPos, i}
		}

		complementsPosition[money-cost] = i
	}

	return []int{}
}

// GetIcecreamChoicesBS Ice cream parlol, gets 2 choices of icecream given certain amount of money
func GetIcecreamChoicesBS(menu []int, money int) []int {
	originalPositions := getOriginalPositionsMatrix(menu)

	sort.Ints(menu)

	for i, cost := range menu {
		complement := money - cost
		complementIndex := binarySearch(menu, complement, i+1, len(menu)-1)

		if complementIndex >= 0 && complementIndex < len(menu) && menu[complementIndex] == complement {
			costIndex := originalPositions[cost][0]
			complementIndex = originalPositions[complement][0]

			if costIndex == complementIndex {
				complementIndex = originalPositions[complement][1]
			}

			if costIndex < complementIndex {
				return []int{costIndex, complementIndex}
			}

			return []int{complementIndex, costIndex}
		}
	}

	return []int{}
}

func binarySearch(menu []int, target, start, end int) int {

	for start <= end {
		mid := int((start + end) / 2)

		if menu[mid] == target {
			return mid
		} else if menu[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func getOriginalPositionsMatrix(menu []int) map[int][]int {
	posMatrix := make(map[int][]int)

	for i, cost := range menu {
		if _, found := posMatrix[cost]; !found {
			posMatrix[cost] = []int{i}
		} else {
			posMatrix[cost] = append(posMatrix[cost], i)
		}
	}

	return posMatrix
}
