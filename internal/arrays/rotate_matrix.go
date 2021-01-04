package arrays

// RotateMatrix90Deg rotates matrix 90 degrees clockwise
func RotateMatrix90Deg(m [][]int) {
	if len(m) <= 1 {
		return
	}

	n := len(m)

	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first

			top := m[first][i]
			// left -> top
			m[first][i] = m[last-offset][first]
			// bottom -> left
			m[last-offset][first] = m[last][last-offset]
			// right -> bottom
			m[last][last-offset] = m[i][last]
			// top -> right
			m[i][last] = top

			// Alternate one liner without temp var (top), but not so readable
			//m[first][i], m[last-offset][first], m[last][last-offset], m[i][last] = m[last-offset][first], m[last][last-offset], m[i][last], m[first][i]
		}
	}
}
