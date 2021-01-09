package stacks

// SortedStack a sorted stack with lower elements on top
type SortedStack struct {
	inStack *regularStack
}

// MakeSortedStack creates sorted stack
func MakeSortedStack() *SortedStack {
	return &SortedStack{
		inStack: new(regularStack),
	}
}

// Peek peeks stacks
func (s *SortedStack) Peek() int {
	return s.inStack.Peek()
}

// Pop pops from stack
func (s *SortedStack) Pop() int {
	return s.inStack.Pop()
}

// Push pushes in sorted order this will be O(n)
func (s *SortedStack) Push(val int) {
	tempStack := new(regularStack)

	for !s.inStack.Empty() && val > s.inStack.Peek() {
		tempStack.Push(s.inStack.Pop())
	}

	s.inStack.Push(val)

	for !tempStack.Empty() {
		s.inStack.Push(tempStack.Pop())
	}
}
