package algorithms

import "fmt"

type Queue struct {
	head *Node
	tail *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) Enqueue(val any) {
	newNode := Node{
		val:  val,
		next: nil,
	}

	q.size += 1

	if q.head == nil && q.tail == nil {
		q.head = &newNode
		q.tail = &newNode
		return
	}

	q.tail.next = &newNode
	q.tail = &newNode
}

func (q *Queue) Dequeue() any {
	if q.head == nil {
		return nil
	}
	q.size -= 1

	node := q.head

	if q.head == q.tail {
		q.head = nil
		q.tail = nil

		return node.val
	}

	q.head = q.head.next
	node.next = nil

	return node.val
}

func (q *Queue) Print() {
	aux := q.head

	for aux != nil {
		fmt.Println("aux:", aux.val)
		aux = aux.next
	}
}

func (q *Queue) GetSize() int {
	return q.size
}
