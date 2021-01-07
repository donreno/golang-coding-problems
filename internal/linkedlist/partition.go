package linkedlist

import (
	s "golang-coding-problems/internal/structs"
)

// Partition partitions list taking all elements lower than p to the left and the others to the right
func Partition(ll *s.LinkedList, p int) *s.LinkedList {
	if ll.Empty() {
		return ll
	}

	var head, tail *s.ListNode

	node := ll.Head

	for node != nil {
		next := node.Next

		if node.Val.(int) < p {
			if head != nil {
				node.Next = head
			}

			head = node
		} else {
			if tail != nil {
				tail.Next = node
			}

			tail = node
		}

		node = next
	}

	tail.Next = nil

	return &s.LinkedList{
		Head: head,
	}
}
