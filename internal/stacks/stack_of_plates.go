package stacks

// StackOfPlates a pile of stacks
type StackOfPlates struct {
	treshold          int
	currentStackIndex int
	stacks            map[int]*regularStack
}

type regularStack struct {
	size int
	top  *regularStackNode
}

type regularStackNode struct {
	val  int
	next *regularStackNode
}

// MakeStackOfPlates makes a stack of plates
func MakeStackOfPlates(treshold int) *StackOfPlates {
	stackOfPlates := new(StackOfPlates)
	stackOfPlates.treshold = treshold
	stackOfPlates.currentStackIndex = 0
	stackOfPlates.stacks = make(map[int]*regularStack)
	stackOfPlates.stacks[0] = &regularStack{
		size: 0,
	}

	return stackOfPlates
}

// Peek stack of plates peek
func (p *StackOfPlates) Peek() int {
	return p.stacks[p.currentStackIndex].Peek()
}

func (p *StackOfPlates) Push(val int) {

	if p.stacks[p.currentStackIndex].size == p.treshold {
		p.currentStackIndex++
		p.stacks[p.currentStackIndex] = &regularStack{
			size: 0,
		}
	}

	p.stacks[p.currentStackIndex].Push(val)
}

func (p *StackOfPlates) Pop() int {

	if p.stacks[p.currentStackIndex].size == 0 && p.currentStackIndex > 0 {
		p.currentStackIndex--
	}

	return p.stacks[p.currentStackIndex].Pop()
}

// base stack

func (r *regularStack) Peek() int {
	if r.top == nil {
		return -1
	}

	return r.top.val
}

func (r *regularStack) Pop() int {
	if r.top == nil {
		return -1
	}

	defer func() {
		r.top = r.top.next
		r.size--
	}()

	return r.top.val
}

func (r *regularStack) Push(val int) {
	node := &regularStackNode{
		val: val,
	}

	if r.top != nil {
		node.next = r.top
	}

	r.size++

	r.top = node
}

func (r *regularStack) Empty() bool {
	return r.top == nil
}
