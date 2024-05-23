package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestBFSExists(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	res := algorithms.BreadthFirstSearch(binTree, 21)
	expected := true

	assert.Equal(t, expected, res)
}

func TestBFSNotExists(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	res := algorithms.BreadthFirstSearch(binTree, 1)
	expected := false

	assert.Equal(t, expected, res)
}
