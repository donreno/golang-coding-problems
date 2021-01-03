package structs

const hashtableFixedSize = 16

// Hashtable interface
type Hashtable interface {
	Put(string, interface{})
	Get(string) interface{}
}

// hashtable struct
type hashtable struct {
	table []*LinkedList
}

// KeyValuePair as valuepair for hashtable
type keyValuePair struct {
	key   string
	value interface{}
}

// MakeHashtable creates a hashtable
func MakeHashtable() Hashtable {
	return &hashtable{
		table: make([]*LinkedList, hashtableFixedSize),
	}
}

// Put puts an element on the hashtable
func (h *hashtable) Put(key string, val interface{}) {
	index := h.hash(key)

	if h.table[index] == nil {
		h.table[index] = new(LinkedList)
	}

	foundElementIndex := -1
	for index, kv := range h.table[index].ToSlice() {
		if kv.(*keyValuePair).key == key {
			foundElementIndex = index
			break
		}
	}

	if foundElementIndex != -1 {
		h.table[index].Remove(foundElementIndex)
	}

	h.table[index].Add(&keyValuePair{
		key:   key,
		value: val,
	})
}

// Get gets an element from hashtable
func (h *hashtable) Get(key string) interface{} {
	index := h.hash(key)

	if h.table[index] == nil {
		return nil
	}

	for _, kv := range h.table[index].ToSlice() {
		if kv.(*keyValuePair).key == key {
			return kv.(*keyValuePair).value
		}
	}

	return nil
}

// hash function that just kind of randomly created to fit into the table size
func (h *hashtable) hash(key string) int {
	sumChars := 0

	i := 0
	for _, r := range key {
		if i > 4 {
			break
		}

		sumChars = sumChars + int(r-'a')
		i++
	}

	index := int(sumChars / hashtableFixedSize)

	if index >= hashtableFixedSize {
		return index % hashtableFixedSize
	}

	return index
}
