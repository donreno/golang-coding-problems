package structs

// Stack Simple stack implementation
type Stack interface {
	Top() interface{}
	Pop() interface{}
	Push(interface{})
	Empty() bool
}

type stack struct {
	top *stackNode
}

type stackNode struct {
	value interface{}
	next  *stackNode
}

// MakeStack creates a new stack instance
func MakeStack() Stack {
	return new(stack)
}

func (s *stack) Top() interface{} {
	if s.Empty() {
		return nil
	}

	return s.top.value
}

func (s *stack) Push(in interface{}) {
	inNode := &stackNode{
		value: in,
		next:  s.top,
	}

	s.top = inNode
}

func (s *stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	defer func() { s.top = s.top.next }()

	return s.top.value
}

func (s *stack) Empty() bool {
	return s.top == nil
}
