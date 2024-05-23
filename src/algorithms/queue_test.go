package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := algorithms.NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	assert.Equal(t, 4, queue.GetSize())
	queue.Dequeue()

	assert.Equal(t, 3, queue.GetSize())

}

func TestQueue2(t *testing.T) {
	queue := algorithms.NewQueue()

	queue.Enqueue(0)
	val := queue.Dequeue()
	assert.Equal(t, 0, val)

	test := queue.Dequeue()
	assert.Equal(t, nil, test)

}

func TestQueue3(t *testing.T) {
	queue := algorithms.NewQueue()

	queue.Enqueue(0)
	val := queue.Dequeue()
	assert.Equal(t, 0, val)

	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Enqueue(0)

	assert.Equal(t, 4, queue.GetSize())
	assert.Equal(t, 3, queue.Dequeue())
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 1, queue.Dequeue())
	assert.Equal(t, 0, queue.Dequeue())

	assert.Equal(t, 0, queue.GetSize())
	assert.Equal(t, nil, queue.Dequeue())
	queue.Enqueue(3)
	queue.Enqueue(2)
	assert.Equal(t, 2, queue.GetSize())
	assert.Equal(t, 3, queue.Dequeue())
	assert.Equal(t, 2, queue.Dequeue())
}
