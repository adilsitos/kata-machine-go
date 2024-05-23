package graphs_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms/graphs"
	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	graph := generateGraph2()

	assert.Equal(t, []int{0, 1, 4, 5, 6}, graphs.DijkstraShortestPath(0, 6, graph))

}
