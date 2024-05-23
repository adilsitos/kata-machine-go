package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	linkedList := algorithms.NewLinkedList()

	linkedList.Append(4)
	linkedList.Append(1)
	linkedList.Append(5)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Prepend(90)
	linkedList.Prepend(40)
	linkedList.Prepend(60)

	assert.Equal(t, 8, linkedList.Length())

	linkedList.TraverseAndPrint()
}

func TestLinkedListWithInsertAt(t *testing.T) {
	linkedList := algorithms.NewLinkedList()

	linkedList.Append(4)
	linkedList.Append(1)
	linkedList.Append(5)
	err := linkedList.InsertAt(100, 4)
	assert.Error(t, err)

	sizeLinkedList := linkedList.Length()
	assert.Equal(t, 3, sizeLinkedList)

	err = linkedList.InsertAt(1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 4, linkedList.Length())
}
