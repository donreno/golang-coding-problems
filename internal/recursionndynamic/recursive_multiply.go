package recursionndynamic

// RecursiveMultiply multiplies a number recursively
func RecursiveMultiply(num, multiplyBy int) int {
	if num < 0 && multiplyBy >= 0 {
		return -recursiveMultiply(-num, multiplyBy)
	}

	if num >= 0 && multiplyBy < 0 {
		return -recursiveMultiply(num, -multiplyBy)
	}

	if num < 0 && multiplyBy < 0 {
		num = -num
		multiplyBy = -multiplyBy
	}

	return recursiveMultiply(num, multiplyBy)
}

func recursiveMultiply(num, multiplyBy int) int {
	if multiplyBy == 0 {
		return 0
	}

	return num + recursiveMultiply(num, multiplyBy-1)
}
