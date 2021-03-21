package linkedlist

import (
	s "golang-coding-problems/internal/structs"
	"math"
)

// MergeKSortedLinkedLists merge k sorted lists
func MergeKSortedLinkedLists(kLists []*s.ListNode) *s.ListNode {
	allNil := false
	head := &s.ListNode{Val: nil}
	current := head

	for !allNil {
		allNil = true
		minIndex, minVal := -1, math.MaxInt32
		for i := range kLists {
			if kLists[i] == nil {
				continue
			}

			allNil = false
			val := kLists[i].Val.(int)
			if val < minVal {
				minIndex, minVal = i, val
			}
		}

		if minIndex != -1 {
			// HEAD node
			if current.Val == nil {
				current.Val = kLists[minIndex].Val
			} else {
				current.Next = &s.ListNode{Val: kLists[minIndex].Val}
				current = current.Next
			}

			kLists[minIndex] = kLists[minIndex].Next
		}
	}

	return head
}
