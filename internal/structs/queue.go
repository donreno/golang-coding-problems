package structs

// Queue represents a Queue
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Head() interface{}
	Empty() bool
}

type queue struct {
	head *queueNode
}

type queueNode struct {
	value interface{}
	next  *queueNode
}

// MakeQueue creates a new Queue
func MakeQueue() Queue {
	return new(queue)
}

func (q *queue) Empty() bool {
	return q.head == nil
}

func (q *queue) Head() interface{} {
	return q.head.value
}

func (q *queue) Enqueue(in interface{}) {
	newNode := &queueNode{
		value: in,
	}

	if q.Empty() {
		q.head = newNode
		return
	}

	currNode := q.head

	for ; currNode.next != nil; currNode = currNode.next {
	}

	currNode.next = newNode
}

func (q *queue) Dequeue() interface{} {
	if q.Empty() {
		return nil
	}

	defer func() {
		q.head = q.head.next
	}()

	return q.Head()
}
