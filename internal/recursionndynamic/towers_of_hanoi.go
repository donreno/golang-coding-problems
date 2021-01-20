package recursionndynamic

import (
	s "golang-coding-problems/internal/structs"
)

// TowersOfHanoi solves towers of hanoi
func TowersOfHanoi(n int, source, target, spare s.Stack) {
	if n == 0 {
		return
	}

	TowersOfHanoi(n-1, source, spare, target)
	target.Push(source.Pop())
	TowersOfHanoi(n-1, spare, target, source)
}
