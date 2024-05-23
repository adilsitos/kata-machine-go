package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestTreeCompareEqual(t *testing.T) {
	a := algorithms.NewBinaryTree()
	a.InsertIntoTree(7)
	a.InsertIntoTree(23)
	a.InsertIntoTree(3)
	a.InsertIntoTree(5)
	a.InsertIntoTree(4)
	a.InsertIntoTree(18)
	a.InsertIntoTree(21)

	b := algorithms.NewBinaryTree()
	b.InsertIntoTree(7)
	b.InsertIntoTree(23)
	b.InsertIntoTree(3)
	b.InsertIntoTree(5)
	b.InsertIntoTree(4)
	b.InsertIntoTree(18)
	b.InsertIntoTree(21)

	assert.Equal(t, true, algorithms.Compare(a.Root, b.Root))
}

func TestTreeCompareValueNotEqual(t *testing.T) {
	a := algorithms.NewBinaryTree()
	a.InsertIntoTree(7)
	a.InsertIntoTree(23)
	a.InsertIntoTree(3)
	a.InsertIntoTree(5)
	a.InsertIntoTree(4)
	a.InsertIntoTree(18)
	a.InsertIntoTree(21)

	b := algorithms.NewBinaryTree()
	b.InsertIntoTree(7)
	b.InsertIntoTree(23)
	b.InsertIntoTree(3)
	b.InsertIntoTree(5)
	b.InsertIntoTree(4)
	b.InsertIntoTree(8)
	b.InsertIntoTree(21)

	assert.Equal(t, false, algorithms.Compare(a.Root, b.Root))
}

func TestTreeCompareEstructureNotEqual(t *testing.T) {
	a := algorithms.NewBinaryTree()
	a.InsertIntoTree(7)
	a.InsertIntoTree(23)
	a.InsertIntoTree(3)
	a.InsertIntoTree(5)
	a.InsertIntoTree(4)
	a.InsertIntoTree(21)
	a.InsertIntoTree(18)

	b := algorithms.NewBinaryTree()
	b.InsertIntoTree(7)
	b.InsertIntoTree(23)
	b.InsertIntoTree(3)
	b.InsertIntoTree(5)
	b.InsertIntoTree(4)
	b.InsertIntoTree(8)
	b.InsertIntoTree(21)

	assert.Equal(t, false, algorithms.Compare(a.Root, b.Root))
}
