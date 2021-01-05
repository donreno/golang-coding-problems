package linkedlist

import (
	s "golang-coding-problems/internal/structs"
)

// SumLists sums 2 linkedLists
func SumLists(l1, l2 *s.LinkedList) *s.LinkedList {
	sum := new(s.LinkedList)

	if l1.Empty() {
		return l2
	}

	if l2.Empty() {
		return l1
	}

	currentl1, currentl2, remainder := l1.Head, l2.Head, 0

	for currentl1 != nil {

		l1Val, l2Val := currentl1.Val.(int), currentl2.Val.(int)

		sumVals := l1Val + l2Val
		sumVals, remainder = remainder+sumVals, 0

		if sumVals >= 10 {
			sumVals -= 10
			remainder = 1
		}

		sum.Add(sumVals)

		currentl1 = currentl1.Next
		currentl2 = currentl2.Next
	}

	return sum
}
