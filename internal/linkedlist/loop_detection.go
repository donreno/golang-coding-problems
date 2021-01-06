package linkedlist

import s "golang-coding-problems/internal/structs"

// HasLoop returns true if there is a loop
func HasLoop(ll *s.LinkedList) bool {
	if ll.Empty() {
		return false
	}

	for fast, slow := ll.Head.Next, ll.Head; fast != nil && fast.Next != nil; slow, fast = slow.Next, fast.Next.Next {
		if fast == slow {
			return true
		}
	}

	return false
}
