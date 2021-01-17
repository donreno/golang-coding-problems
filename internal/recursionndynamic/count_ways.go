package recursionndynamic

// CountWays counts the ways to go trough the stairs
func CountWays(n int) int {
	return countWays(n, make([]int, n+1))
}

func countWays(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}

	if memo[n] == 0 {
		memo[n] = countWays(n-1, memo) + countWays(n-2, memo) + countWays(n-3, memo)
	}

	return memo[n]
}
