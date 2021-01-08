package stacks

import "errors"

const FixedStackSize = 3

// ThreeStacksInOneArray struct that holds three stacks in one array
type ThreeStacksInOneArray struct {
	numberOfStacks int
	stackCapacity  int
	values         []int
	sizes          []int
}

// MakeThreeInOneStack creates a three in one (fixed size stack)
func MakeThreeInOneStack(size int) *ThreeStacksInOneArray {
	return &ThreeStacksInOneArray{
		numberOfStacks: FixedStackSize,
		stackCapacity:  size,
		values:         make([]int, size*FixedStackSize),
		sizes:          make([]int, FixedStackSize),
	}
}

// Push pushes to stack, stackNum should be numbers between 0 and FixedStackSize -1
func (t *ThreeStacksInOneArray) Push(val, stackNum int) error {
	if stackNum < 0 || stackNum >= FixedStackSize {
		return errors.New("Bad stack number")
	}

	if t.sizes[stackNum] == t.stackCapacity {
		return errors.New("Stack is already full")
	}

	t.sizes[stackNum]++
	t.values[t.getStackTopIndex(stackNum)] = val

	return nil
}

// Peek peek from chosen stack
func (t *ThreeStacksInOneArray) Peek(stackNum int) (int, error) {
	if stackNum < 0 || stackNum >= FixedStackSize {
		return -1, errors.New("Bad stack number")
	}

	if t.Empty(stackNum) {
		return -1, errors.New("Stack is empty")
	}

	return t.values[t.getStackTopIndex(stackNum)], nil
}

// Pop pops an element from specific stack
func (t *ThreeStacksInOneArray) Pop(stackNum int) (int, error) {
	if stackNum < 0 || stackNum >= FixedStackSize {
		return -1, errors.New("Bad stack number")
	}

	if t.Empty(stackNum) {
		return -1, errors.New("Stack is empty")
	}

	defer func() {
		t.sizes[stackNum]--
	}()

	return t.values[t.getStackTopIndex(stackNum)], nil
}

// Empty checks if stack is empty
func (t *ThreeStacksInOneArray) Empty(stackNum int) bool {
	if stackNum < 0 || stackNum >= FixedStackSize {
		return true
	}

	return t.sizes[stackNum] == 0
}

func (t *ThreeStacksInOneArray) getStackTopIndex(stackNum int) int {

	return FixedStackSize*stackNum + t.sizes[stackNum] - 1
}

// Min finds minimum value on stack
func (t *ThreeStacksInOneArray) Min(stackNum int) (int, error) {
	if stackNum < 0 || stackNum >= FixedStackSize {
		return -1, errors.New("Bad stack number")
	}

	if t.Empty(stackNum) {
		return -1, errors.New("Stack is empty")
	}

	i := FixedStackSize * stackNum

	min := t.values[i]

	for i = i + 1; i <= t.getStackTopIndex(stackNum); i++ {
		if t.values[i] < min {
			min = t.values[i]
		}
	}

	return min, nil
}
