package lru

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

func NewNode(value any) *Node {
	return &Node{Value: value, next: nil, prev: nil}
}

type LRU struct {
	length        int
	head          *Node
	tail          *Node
	lookup        map[any]*Node
	reverseLookup map[*Node]any // the reverse lookup is used to get the key when a node needs be trimmed
	capacity      int
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		length:        0,
		head:          nil,
		tail:          nil,
		lookup:        make(map[any]*Node),
		reverseLookup: make(map[*Node]any),
		capacity:      capacity,
	}
}

func (l *LRU) Update(key any, value any) {
	// does it exist?
	// get(key)

	// if it doesn't exist -> insert
	// 		- check capacity and evict if over
	// it it does, we need to update to the front of the list
	// and  update the value
	// var node *Node
	node, ok := l.lookup[key]
	if !ok {
		node := NewNode(value)
		l.length++
		l.prepend(node)
		l.trimCache()
		l.lookup[key] = node
		l.reverseLookup[node] = key
	} else {
		l.detach(node)
		l.prepend(node)
		node.Value = value
	}

}

func (l *LRU) Get(key any) any {
	// check the cache for existence

	// update the value we found and move it to the front
	// it becomes the most recent used (newest possible)

	// return out the value found, or nil

	node, ok := l.lookup[key]
	if !ok {
		return nil
	}

	l.detach(node)
	l.prepend(node)

	return node.Value
}

func (l *LRU) detach(node *Node) {

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}

	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *LRU) prepend(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRU) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail

	l.detach(l.tail)

	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.length--
}
