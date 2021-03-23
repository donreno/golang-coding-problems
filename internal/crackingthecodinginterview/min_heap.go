package crackingthecodinginterview

// MinHeap a min heap implementation
type MinHeap struct {
	arr      []int
	size     int
	capacity int
}

// BuildMinHeap builds a min heap
func BuildMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		arr:      make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func getParentIndex(childIndex int) int {
	return int((childIndex - 1) / 2)
}

func (h *MinHeap) hasLeftChild(index int) bool {
	return getLeftChildIndex(index) < len(h.arr)
}

func (h *MinHeap) hasRightChild(index int) bool {
	return getRightChildIndex(index) < len(h.arr)
}

func hasParent(index int) bool {
	return getParentIndex(index) >= 0
}

func (h *MinHeap) leftChild(index int) int {
	return h.arr[getLeftChildIndex(index)]
}

func (h *MinHeap) rightChild(index int) int {
	return h.arr[getRightChildIndex(index)]
}

func (h *MinHeap) parent(index int) int {
	return h.arr[getParentIndex(index)]
}

func (h *MinHeap) swap(firstIndex, secondIndex int) {
	h.arr[firstIndex], h.arr[secondIndex] = h.arr[secondIndex], h.arr[firstIndex]
}

func (h *MinHeap) ensureExtraCapacity() {
	if h.size == h.capacity {
		h.capacity *= 2
		newArr := make([]int, h.capacity)
		copy(newArr, h.arr)
		h.arr = newArr
	}
}

// Poll polls top element of heap
func (h *MinHeap) Poll() int {
	if h.size == 0 {
		return -1
	}

	item := h.arr[0]
	h.arr = h.arr[1:]
	h.size--

	h.heapifyDown()

	return item
}

// Add adds an item to heap
func (h *MinHeap) Add(item int) {
	h.ensureExtraCapacity()
	h.arr[h.size] = item
	h.size++
	h.heapifyUp()
}

func (h *MinHeap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := getLeftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIndex = getRightChildIndex(index)
		}

		if h.arr[index] < h.arr[smallerChildIndex] {
			break
		} else {
			h.swap(index, smallerChildIndex)
			index = smallerChildIndex
		}
	}
}

func (h *MinHeap) heapifyUp() {
	index := h.size - 1
	for hasParent(index) && h.parent(index) > h.arr[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}
