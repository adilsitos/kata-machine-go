package algorithms_test

import (
	"slices"
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	initArr := []int{5, 4, 3, 20, 12, 11, 9, 10}
	element := 20

	res := algorithms.BinarySearch(initArr, element)
	slices.Sort(initArr)

	assert.Equal(t, element, initArr[res])
}

func TestBinarySearchValueDoesNotExist(t *testing.T) {
	initArr := []int{5, 4, 3, 20, 12, 11, 9, 10}
	element := 400

	res := algorithms.BinarySearch(initArr, element)
	slices.Sort(initArr)

	assert.Equal(t, -1, res)
}
