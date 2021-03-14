package googleinterview

// SpecialLL special linkedlist with arbitrary pointer
type SpecialLL struct {
	Data      int
	Next      *SpecialLL
	Arbitrary *SpecialLL
}

// DeepCopyLinkedList deep copies linkedlist with arbitrary pointer
func DeepCopyLinkedList(ll *SpecialLL) *SpecialLL {
	visited := make(map[int]*SpecialLL)
	dcpLL := new(SpecialLL)

	deepCopyLL(ll, dcpLL, visited)

	return dcpLL
}

func deepCopyLL(ll, dcpLL *SpecialLL, visited map[int]*SpecialLL) {
	if memo, found := visited[ll.Data]; found {
		dcpLL.Data = memo.Data
		dcpLL.Next = memo.Next
		dcpLL.Arbitrary = memo.Arbitrary
		return
	}

	dcpLL.Data = ll.Data
	visited[ll.Data] = dcpLL

	if ll.Next != nil {
		dcpLL.Next = new(SpecialLL)
		deepCopyLL(ll.Next, dcpLL.Next, visited)
	}

	if ll.Arbitrary != nil {
		dcpLL.Arbitrary = new(SpecialLL)
		deepCopyLL(ll.Arbitrary, dcpLL.Arbitrary, visited)
	}
}
