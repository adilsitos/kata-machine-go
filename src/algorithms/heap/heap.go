package heap

import (
	"fmt"
	"math"
)

// min heap
// at a node, all child nodes must be bigger than it
type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Delete() int {
	if len(h.arr) == 0 {
		return -1
	}

	arrSize := len(h.arr) - 1

	out := h.arr[0]

	lastElement := h.arr[arrSize]
	h.arr[0] = lastElement
	h.arr = h.arr[:len(h.arr)-1]

	h.heapfyDown(0)

	return out
}

func (h *Heap) Insert(val int) {
	h.arr = append(h.arr, val)
	arrSize := len(h.arr) - 1

	h.heapfyUp(arrSize)
}

func (h *Heap) heapfyDown(idx int) {
	leftPos := h.leftChild(idx)
	rightPos := h.rightChild(idx)

	arrSize := len(h.arr)

	if idx >= arrSize || leftPos >= arrSize {
		return
	}

	var leftVal = math.MaxInt
	var rightVal = math.MaxInt

	if leftPos < arrSize {
		leftVal = h.arr[leftPos]
	}

	if rightPos < arrSize {
		rightVal = h.arr[rightPos]
	}

	element := h.arr[idx]

	if leftVal < rightVal && element > leftVal {
		h.arr[idx] = leftVal
		h.arr[leftPos] = element
		h.heapfyDown(leftPos)
	}

	if rightVal < leftVal && element > rightVal {
		h.arr[idx] = rightVal
		h.arr[rightPos] = element
		h.heapfyDown(rightPos)
	}
}

func (h *Heap) heapfyUp(idx int) {
	if idx < 0 {
		return
	}

	element := h.arr[idx]
	parentPos := h.parent(idx)
	parentVal := h.arr[parentPos]

	if element < parentVal {
		//swap
		h.arr[parentPos] = element
		h.arr[idx] = parentVal

		h.heapfyUp(parentPos)
	}
}

func (h *Heap) leftChild(idx int) int {
	return 2*idx + 1
}

func (h *Heap) rightChild(idx int) int {
	return 2*idx + 2
}

func (h *Heap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap) Size() int {
	return len(h.arr)
}

func (h *Heap) Print() {
	fmt.Println(h.arr)
}
