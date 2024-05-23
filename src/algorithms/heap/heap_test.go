package heap_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms/heap"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := heap.NewHeap()

	h.Insert(3)
	h.Insert(400)
	h.Insert(30)
	h.Insert(20)
	h.Insert(10)
	h.Insert(90)
	h.Insert(11)
	h.Insert(230)
	h.Insert(120)
	h.Insert(1)

	h.Print()

	assert.Equal(t, 10, h.Size())

	assert.Equal(t, 1, h.Delete())
	assert.Equal(t, 3, h.Delete())
	assert.Equal(t, 10, h.Delete())
	assert.Equal(t, 11, h.Delete())
	assert.Equal(t, 20, h.Delete())
	assert.Equal(t, 30, h.Delete())
	assert.Equal(t, 90, h.Delete())
	assert.Equal(t, 120, h.Delete())
	assert.Equal(t, 230, h.Delete())
	assert.Equal(t, 400, h.Delete())

	assert.Equal(t, 0, h.Size())
}
