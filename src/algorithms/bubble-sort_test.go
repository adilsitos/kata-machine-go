package algorithms_test

import (
	"slices"
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	initialArr := []int{4, 10, 32, 1, 2, 5, 12, 4123, 44, 5, 1, 3}

	slices.Sort(initialArr)

	actualArr := algorithms.BubbleSort([]int{4, 10, 32, 1, 2, 5, 12, 4123, 44, 5, 1, 3})

	assert.Equal(t, initialArr, actualArr)

}
