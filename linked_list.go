package golang_coding_problems

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

type ListNode struct {
	val  interface{}
	next *ListNode
}

func (l *LinkedList) Empty() bool {
	if l.head == nil {
		return true
	}

	return false
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Head() interface{} {
	if l.Empty() {
		return nil
	}

	return l.head.val
}

func (l *LinkedList) Tail() interface{} {
	if l.Empty() {
		return nil
	}

	return l.tail.val
}

func (l *LinkedList) Add(in interface{}) {
	newNode := &ListNode{
		val: in,
	}

	if l.Empty() {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next, l.tail = newNode, newNode
	}

	l.size++
}

func (l *LinkedList) Remove(index int) {
	if l.Empty() {
		return
	}

	defer func() {
		l.size--
	}()

	if index == 0 {
		if l.size == 1 {
			l.head, l.tail = nil, nil
		} else {
			l.head = l.head.next
		}
	} else {
		prev, current := l.head, l.head.next

		for i := 1; current != nil; prev, current, i = current, current.next, i+1 {
			if i == index {
				prev.next = current.next
				_ = current
				break
			}
		}
	}
}

func (l *LinkedList) Get(index int) interface{} {
	if l.Empty() {
		return nil
	}

	currNode := l.head

	for i := 0; currNode != nil; currNode, i = currNode.next, i+1 {
		if i == index {
			return currNode.val
		}
	}

	return nil
}

func (l *LinkedList) ToSlice() (sliceList []interface{}) {
	sliceList = make([]interface{}, l.size)

	if l.Empty() {
		return
	}

	currNode := l.head

	for i := 0; currNode != nil; currNode, i = currNode.next, i+1 {
		sliceList[i] = currNode.val
	}

	return
}
