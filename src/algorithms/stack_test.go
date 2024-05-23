package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := algorithms.NewStack()

	val := stack.Pop()
	assert.Equal(t, nil, val)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	assert.Equal(t, 4, stack.GetSize())

	val = stack.Pop()
	assert.Equal(t, 4, val)

	val = stack.Pop()
	assert.Equal(t, 3, val)

	assert.Equal(t, 2, stack.GetSize())
}
