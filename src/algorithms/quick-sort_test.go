package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	arr := []int{9, 8, 7, 6, 4, 5}

	algorithms.Partition(arr, 0, len(arr)-1)

	expectedRes := []int{4, 5, 7, 6, 5, 8}

	assert.Equal(t, expectedRes, arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 4, 5}

	algorithms.QuickSort(arr)

	expectedRes := []int{4, 5, 6, 7, 8, 9}

	assert.Equal(t, expectedRes, arr)
}
