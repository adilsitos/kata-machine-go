package algorithms

import (
	"errors"
	"fmt"
)

type Node struct {
	val  any
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	idx  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(value any) {
	newNode := Node{val: value, next: nil}
	l.idx++

	if l.head == nil && l.tail == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	if l.head == nil {
		panic("head is nil")
	}

	l.tail.next = &newNode
	l.tail = &newNode
}

func (l *LinkedList) Prepend(value any) {
	newNode := Node{val: value, next: nil}
	l.idx++

	if l.head == nil && l.tail == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	newNode.next = l.head
	l.head = &newNode
}

func (l *LinkedList) InsertAt(idx int, newVal any) error {
	if idx > l.idx || idx < 0 {
		return errors.New("idx is invalid")
	}

	newNode := Node{
		val:  newVal,
		next: nil,
	}

	aux := l.head
	if l.idx == 0 && l.head == nil && l.tail == nil {
		l.head = &newNode
		l.tail = &newNode
		return nil
	}

	var prev *Node
	idxAux := 0

	for aux != nil {
		if idxAux == idx {
			if prev != nil {
				// there is only one node on the linked list
				prev.next = &newNode
			}

			newNode.next = aux
			l.idx++

			return nil
		}

		idxAux++
		prev = aux
		aux = aux.next
	}

	return nil
}

func (l *LinkedList) RemoveAt(idx int) error {
	if idx > l.idx || idx < 0 {
		return errors.New("idx is invalid")
	}

	if l.head == nil || l.tail == nil {
		return nil
	}

	var prev *Node
	aux := l.head
	auxIdx := 0

	for aux != nil {
		if auxIdx == idx {
			if prev != nil {
				prev.next = aux.next
				// free aux
				return nil
			}

			// first element
			l.head = aux.next
			aux.next = nil

			return nil
		}

		prev = aux
		aux = aux.next
	}

	return nil
}

func (l *LinkedList) TraverseAndPrint() {
	aux := l.head

	for aux != nil {
		fmt.Println(aux.val)
		aux = aux.next
	}
}

func (l *LinkedList) Length() int {
	return l.idx
}
