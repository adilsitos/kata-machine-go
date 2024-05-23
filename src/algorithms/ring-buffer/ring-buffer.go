package main

import "fmt"

type Ring struct {
	capacity int
	size     int
	head     int
	tail     int
	buffer   []int
}

const defaultCap = 10

func NewRing(capacity int) *Ring {
	if capacity <= 0 {
		capacity = defaultCap
	}

	return &Ring{
		capacity: capacity,
		buffer:   make([]int, capacity),
	}
}

func (ring *Ring) Push(item int) {
	// the tail encounters the head
	if ring.tail == ring.head && ring.size == ring.capacity {
		ring.head = (ring.head + 1) % ring.capacity
	}

	ring.buffer[ring.tail] = item
	ring.tail = (ring.tail + 1) % ring.capacity

	if ring.size < ring.capacity {
		ring.size++
	}
}

func (ring *Ring) Pop() int {
	if ring.size < 1 {
		return -1
	}

	item := ring.buffer[ring.head]
	ring.buffer[ring.head] = -1
	ring.size--

	ring.head = (ring.head + 1) % ring.capacity

	return item
}

func (ring *Ring) Print() {
	if ring.size == 0 {
		fmt.Println("empty")
		return
	}

	fmt.Printf("[%d", ring.buffer[ring.head])
	i := (ring.head + 1) % ring.capacity

	for i != ring.tail {
		fmt.Printf(", %d", ring.buffer[i])

		i = (i + 1) % ring.capacity
	}

	fmt.Println("]")
}
