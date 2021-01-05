package linkedlist

import (
	s "golang-coding-problems/internal/structs"
)

// RemoveDups removes duplicates from linkedlist (Assumes element is an int)
func RemoveDups(ll *s.LinkedList) {
	if ll == nil || ll.Empty() {
		return
	}

	foundElems := make(map[int]bool)

	current := ll.Head
	var prev *s.ListNode

	for current != nil {
		val := current.Val.(int)

		if foundElems[val] {
			prev.Next, current = current.Next, current.Next
		} else {
			foundElems[val] = true
			prev, current = current, current.Next
		}
	}
}
