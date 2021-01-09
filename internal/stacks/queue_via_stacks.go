package stacks

// QueueViaStack a queue via stack
type QueueViaStack struct {
	stackQ *regularStack
}

// MakeQueueViaStack makes queue via stack
func MakeQueueViaStack() *QueueViaStack {
	return &QueueViaStack{
		stackQ: new(regularStack),
	}
}

// Head head of queue
func (q *QueueViaStack) Head() int {
	return q.stackQ.Peek()
}

// Dequeue returns head of the queue and dequeues
func (q *QueueViaStack) Dequeue() int {
	return q.stackQ.Pop()
}

// Enqueue enqueues on stack
func (q *QueueViaStack) Enqueue(val int) {
	tempStack := new(regularStack)

	for !q.stackQ.Empty() {
		tempStack.Push(q.stackQ.Pop())
	}

	q.stackQ.Push(val)

	for !tempStack.Empty() {
		q.stackQ.Push(tempStack.Pop())
	}
}
