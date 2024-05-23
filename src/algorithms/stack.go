package algorithms

import "fmt"

type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
}

func (s *Stack) Push(value any) {
	s.size++
	newNode := Node{
		val:  value,
		next: nil,
	}

	if s.top == nil {
		s.top = &newNode
	}

	newNode.next = s.top

	s.top = &newNode
}

func (s *Stack) Pop() any {
	if s.top == nil {
		return nil
	}

	node := s.top

	s.top = s.top.next

	node.next = nil
	// free node

	s.size--

	return node.val
}

func (s *Stack) Print() {
	aux := s.top

	for s.top != nil {
		fmt.Println(aux.val)

		aux = aux.next
	}
}

func (s *Stack) GetSize() int {
	return s.size
}
