package graphs_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms/graphs"
	"github.com/stretchr/testify/assert"
)

var matrix = graphs.GraphAdjacencyMatrix{
	{0, 3, 1, 0, 0, 0, 0},  // 0
	{0, 0, 0, 0, 1, 0, 0},  // 1
	{0, 0, 0, 7, 0, 0, 0},  // 2
	{0, 0, 0, 0, 0, 0, 0},  // 3
	{0, 1, 0, 5, 0, 2, 0},  // 4
	{0, 0, 18, 0, 0, 0, 1}, // 5
	{0, 0, 0, 1, 0, 0, 1},  // 6
}

func TestBFSOnAdjacencyMatrix(t *testing.T) {
	expected := []int{0, 1, 4, 5, 6}

	assert.Equal(t, expected, graphs.BFSOnAdjancyMatrix(matrix, 0, 6))
	assert.Equal(t, []int{}, graphs.BFSOnAdjancyMatrix(matrix, 6, 0))
	assert.Equal(t, []int{}, graphs.BFSOnAdjancyMatrix(matrix, 0, 0))
}
