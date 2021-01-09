package stacks

// MinStack Stack with min element so push, pop and min funcs are all O(1)
type MinStack struct {
	min         int
	previousMin int
	top         *MinStackNode
}

// MinStackNode a min stack node
type MinStackNode struct {
	Val  int
	Next *MinStackNode
}

// Peek returns top of the stack without popping it
func (m *MinStack) Peek() int {
	if m.top == nil {
		return -1
	}

	return m.top.Val
}

// Push pushes a new element
func (m *MinStack) Push(val int) {
	m.updateMin(val, m.min)

	node := &MinStackNode{
		Val: val,
	}

	defer func() {
		m.top = node
	}()

	if m.top != nil {
		node.Next = m.top
	}
}

// Pop pops top element of the stack
func (m *MinStack) Pop() int {
	if m.top == nil {
		return -1
	}

	defer func() {
		m.top = m.top.Next
	}()

	if m.top.Val == m.min {
		m.min = m.previousMin
	}

	return m.top.Val
}

// Min returns min element
func (m *MinStack) Min() int {
	return m.min
}

func (m *MinStack) updateMin(val, min int) {
	if m.top == nil || val < min {
		m.previousMin, m.min = m.min, val
	}
}
