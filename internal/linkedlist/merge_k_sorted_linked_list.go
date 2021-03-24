package linkedlist

import (
	s "golang-coding-problems/internal/structs"
)

// MergeKSortedLinkedLists merge k sorted lists
func MergeKSortedLinkedLists(kLists []*s.ListNode) *s.ListNode {
	start, end := 0, len(kLists)-1

	for start < end {
		mergeSortedLinkedLists(kLists, start, end)
		end--
	}

	return kLists[start]
}

func mergeSortedLinkedLists(kLists []*s.ListNode, start, end int) {
	firstCurrent, secondCurrent := kLists[start], kLists[end]
	var head *s.ListNode
	var current *s.ListNode

	if firstCurrent.Val.(int) < secondCurrent.Val.(int) {
		head = s.MakeListNode(firstCurrent.Val.(int))
		firstCurrent = firstCurrent.Next
	} else {
		head = s.MakeListNode(secondCurrent.Val.(int))
		secondCurrent = secondCurrent.Next
	}

	current = head

	for firstCurrent != nil && secondCurrent != nil {
		first, second := firstCurrent.Val.(int), secondCurrent.Val.(int)
		var newNode *s.ListNode

		if first < second {
			newNode = s.MakeListNode(first)
			firstCurrent = firstCurrent.Next
		} else {
			newNode = s.MakeListNode(second)
			secondCurrent = secondCurrent.Next
		}

		current.Next = newNode
		current = current.Next
	}

	for firstCurrent != nil {
		node := s.MakeListNode(firstCurrent.Val)
		current.Next = node
		current = current.Next
		firstCurrent = firstCurrent.Next
	}

	for secondCurrent != nil {
		node := s.MakeListNode(secondCurrent.Val)
		current.Next = node
		current = current.Next
		secondCurrent = secondCurrent.Next
	}

	kLists[start] = head
}
