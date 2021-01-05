package structs

// LinkedList contains head, tail and size
type LinkedList struct {
	Head *ListNode
	Tail *ListNode
	size int
}

// ListNode a node for the linkedlist
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// Empty returns true if empty
func (l *LinkedList) Empty() bool {
	if l.Head == nil {
		return true
	}

	return false
}

// Size returns the size of the list
func (l *LinkedList) Size() int {
	return l.size
}

// Add or AppendToTail
func (l *LinkedList) Add(in interface{}) *LinkedList {
	newNode := &ListNode{
		Val: in,
	}

	if l.Empty() {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next, l.Tail = newNode, newNode
	}

	l.size++

	return l
}

// Remove removes Ith element from Linkedlist
func (l *LinkedList) Remove(index int) *LinkedList {
	if l.Empty() {
		return l
	}

	defer func() {
		l.size--
	}()

	if index == 0 {
		if l.size == 1 {
			l.Head, l.Tail = nil, nil
		} else {
			l.Head = l.Head.Next
		}
	} else {
		prev, current := l.Head, l.Head.Next

		for i := 1; current != nil; prev, current, i = current, current.Next, i+1 {
			if i == index {
				prev.Next = current.Next
				_ = current
				break
			}
		}
	}

	return l
}

// Get gets the Ith element from the list
func (l *LinkedList) Get(index int) interface{} {
	if l.Empty() {
		return nil
	}

	currNode := l.Head

	for i := 0; currNode != nil; currNode, i = currNode.Next, i+1 {
		if i == index {
			return currNode.Val
		}
	}

	return nil
}

// ToSlice will create a slice from the Linkedlist
func (l *LinkedList) ToSlice() (sliceList []interface{}) {
	sliceList = make([]interface{}, l.size)

	if l.Empty() {
		return
	}

	currNode := l.Head

	for i := 0; currNode != nil; currNode, i = currNode.Next, i+1 {
		sliceList[i] = currNode.Val
	}

	return
}
