package crackingthecodinginterview

// LinkedListNode linked list node
type LinkedListNode struct {
	Next *LinkedListNode
	Data int
}

// LinkedList wrapper for list nodes
type LinkedList struct {
	Head *LinkedListNode
}

// MakeLinkedListNode makes a new linkedlist node
func MakeLinkedListNode(data int) *LinkedListNode {
	return &LinkedListNode{
		Data: data,
	}
}

// Append appends a new element on linked list
func (ll *LinkedList) Append(data int) {
	if ll.Head == nil {
		ll.Head = MakeLinkedListNode(data)
		return
	}

	current := ll.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = MakeLinkedListNode(data)
}

// Prepend prepends data to linked list
func (ll *LinkedList) Prepend(data int) {
	newHead := MakeLinkedListNode(data)
	newHead.Next = ll.Head
	ll.Head = newHead
}

// DeleteWithValue deletes an element with specific value from linked list
func (ll *LinkedList) DeleteWithValue(data int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
