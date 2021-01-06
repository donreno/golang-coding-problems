package linkedlist

import (
	s "golang-coding-problems/internal/structs"
)

// IsPalindromeFastSlow checks if linkedlist is palindrome
func IsPalindromeFastSlow(ll *s.LinkedList) bool {
	if ll.Empty() {
		return false
	}

	slow, fast := ll.Head, ll.Head
	stack := s.MakeStack()

	// b - u - s - a - s - u - b

	for fast != nil && fast.Next != nil {
		stack.Push(slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		top := stack.Pop()

		if slow.Val != top {
			return false
		}

		slow = slow.Next
	}

	return true
}
