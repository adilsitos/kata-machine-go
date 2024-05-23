package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestDFSExists(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	assert.Equal(t, true, algorithms.DFS(binTree, 18))
}

func TestDFSNotExists(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	assert.Equal(t, false, algorithms.DFS(binTree, 8))
}
