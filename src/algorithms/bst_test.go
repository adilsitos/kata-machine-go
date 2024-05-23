package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	bst := algorithms.NewBST()

	bst.Insert(4)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(2)

	assert.Equal(t, true, bst.Root.Find(2))
	assert.Equal(t, false, bst.Root.Find(1))

	bst.Print()
	bst.Deletion(2)

	assert.Equal(t, false, bst.Root.Find(2))
	bst.Print()
}
