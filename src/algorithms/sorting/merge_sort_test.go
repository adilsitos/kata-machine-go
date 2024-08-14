package sorting_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms/sorting"
	"github.com/stretchr/testify/require"
)

func TestMergeSort(t *testing.T) {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := sorting.MergeSort(unsorted)

	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, sorted)

}
