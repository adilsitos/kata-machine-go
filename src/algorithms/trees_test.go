package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

// func TestPreOrder(t *testing.T) {
// 	binTree := algorithms.BinaryNode{
// 		Value: 3, Right: &algorithms.BinaryNode{
// 			Value: ,

// 		} }

// }

func TestPreOrder(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	res := algorithms.PreOrder(binTree)
	expectedRes := []int{7, 23, 5, 4, 3, 18, 21}
	assert.Equal(t, expectedRes, res)
}

func TestInOder(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	res := algorithms.InOrder(binTree)
	expectedRes := []int{5, 23, 4, 7, 18, 3, 21}
	assert.Equal(t, expectedRes, res)
}

func TestPostOrder(t *testing.T) {
	binTree := algorithms.NewBinaryTree()

	binTree.InsertIntoTree(7)
	binTree.InsertIntoTree(23)
	binTree.InsertIntoTree(3)
	binTree.InsertIntoTree(5)
	binTree.InsertIntoTree(4)
	binTree.InsertIntoTree(18)
	binTree.InsertIntoTree(21)

	res := algorithms.PostOrder(binTree)
	expectedRes := []int{5, 4, 23, 18, 21, 3, 7}
	assert.Equal(t, expectedRes, res)

}
