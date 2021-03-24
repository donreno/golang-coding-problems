package crackingthecodinginterview

// CountPaths Counts paths given I take 1, 2 or 3 steps at a time
func CountPaths(steps int) int {
	paths := []int{0, 1, 2}

	if steps <= 2 {
		return paths[steps]
	}

	for i := 3; i <= steps; i++ {
		count := paths[2] + paths[1] + paths[0]
		paths[0], paths[1], paths[2] = paths[1], paths[2], count
	}

	return paths[2]
}
