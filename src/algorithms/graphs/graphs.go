package graphs

type GraphEdge struct {
	To     int
	Weight int
}

type GraphAdjacencyList [][]GraphEdge

type GraphAdjacencyMatrix [][]int
