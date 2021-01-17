package recursionndynamic

// MemoFibonacci returns the nth term of the fibonacci series
func MemoFibonacci(n int) int {
	return fibonacci(n, make([]int, n+1))
}

func fibonacci(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}

	if memo[n] == 0 {
		memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	}

	return memo[n]
}

// BottomupDPFibonacci bottom up approach for fibonacci
func BottomupDPFibonacci(n int) int {
	if n == 0 {
		return 0
	}

	nMinus1 := 1
	nMinus2 := 0

	for i := 2; i < n; i++ {
		nMinus2, nMinus1 = nMinus1, nMinus1+nMinus2
	}

	return nMinus1 + nMinus2
}
