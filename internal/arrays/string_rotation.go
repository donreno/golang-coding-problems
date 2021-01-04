package arrays

import "strings"

func IsRotation(a, b string) bool {
	if len(a) != len(b) || len(a) == 0 {
		return false
	}

	aTwice := a + a

	return strings.Contains(aTwice, b)
}
