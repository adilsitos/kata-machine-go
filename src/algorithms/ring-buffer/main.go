package main

import "fmt"

func main() {
	ring := NewRing(4)

	for {
		var i int
		fmt.Println("Type a number to insert into the ring buffer")
		fmt.Scan(&i)
		ring.Push(i)
		ring.Print()
	}
}
