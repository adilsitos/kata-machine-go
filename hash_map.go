package katamachinegolang

import "errors"

// hash map using chaining
// keys are only integers

const tableSize = 10

type Node struct {
	key  int
	val  interface{}
	next *Node
}

type HashMap struct {
	arr  []*Node
	size int
}

func NewHashMap() *HashMap {
	return &HashMap{
		arr:  make([]*Node, tableSize),
		size: tableSize,
	}

}

func (h *HashMap) hashing(key int) int {
	return key % tableSize
}

func (h *HashMap) Set(key int, val interface{}) {
	newNode := Node{key: key, val: val, next: nil}

	idx := h.hashing(key)

	curr := h.arr[idx]
	if curr == nil {
		// curr is nil
		h.arr[idx] = &newNode
		return
	}

	// chaining
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &newNode
}

func (h *HashMap) Get(key int) (interface{}, error) {
	idx := h.hashing(key)

	curr := h.arr[idx]

	for curr != nil {
		if curr.key == key {
			return curr.val, nil
		}
		curr = curr.next
	}

	return nil, errors.New("could not find key")
}

func (h *HashMap) Remove(key int) error {
	idx := h.hashing(key)

	prev := &Node{}
	curr := h.arr[idx]

	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			curr.next = nil

			return nil
		}

		prev = curr
		curr = curr.next
	}

	return errors.New("could not find key")
}
