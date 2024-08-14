package graphs_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms/graphs"
	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {
	graph := generateGraph()
	res := graphs.DFSAdjancyList(graph, 0, 6)
	expectedRes := []int{0, 1, 4, 5, 6}

	assert.Equal(t, expectedRes, res)

	// assert.Equal(t, []int{}, graphs.DFSAdjancyList(graph, 6, 0))
}

func generateGraph() graphs.GraphAdjacencyList {
	graph := make(graphs.GraphAdjacencyList, 7)
	graph[0] = []graphs.GraphEdge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}

	// Vertex 1
	graph[1] = []graphs.GraphEdge{
		// {To: 0, Weight: 3},
		// {To: 2, Weight: 4},
		{To: 4, Weight: 1},
	}

	// Vertex 2
	graph[2] = []graphs.GraphEdge{
		// {To: 1, Weight: 4},
		{To: 3, Weight: 7},
		// {To: 0, Weight: 1},
	}

	// Vertex 3 (no outgoing edges)
	graph[3] = []graphs.GraphEdge{
		// {To: 2, Weight: 7},
		// {To: 4, Weight: 5},
		// {To: 6, Weight: 1},
	}

	// Vertex 4
	graph[4] = []graphs.GraphEdge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}

	// Vertex 5
	graph[5] = []graphs.GraphEdge{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
		// {To: 4, Weight: 2},
	}

	// Vertex 6
	graph[6] = []graphs.GraphEdge{
		{To: 3, Weight: 1},
		// {To: 5, Weight: 1},
	}

	return graph
}

func generateGraph2() graphs.GraphAdjacencyList {
	graph := make(graphs.GraphAdjacencyList, 7)
	graph[0] = []graphs.GraphEdge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}

	// Vertex 1
	graph[1] = []graphs.GraphEdge{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	}

	// Vertex 2
	graph[2] = []graphs.GraphEdge{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	}

	// Vertex 3 (no outgoing edges)
	graph[3] = []graphs.GraphEdge{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	}

	// Vertex 4
	graph[4] = []graphs.GraphEdge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}

	// Vertex 5
	graph[5] = []graphs.GraphEdge{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	}

	// Vertex 6
	graph[6] = []graphs.GraphEdge{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	}

	return graph
}
