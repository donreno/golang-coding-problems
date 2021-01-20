package recursionndynamic

// HasMagicindex checks if array has magic index
func HasMagicindex(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	midPoint := len(arr) / 2
	sum := 1

	if arr[midPoint] > midPoint {
		sum = -1
	}

	for i := midPoint; i < len(arr) && i >= 0; i = i + sum {
		val := arr[i]

		if val == i {
			return true
		}

		if val > 0 && val < len(arr) && arr[val] == val {
			return true
		}
	}

	return false
}
