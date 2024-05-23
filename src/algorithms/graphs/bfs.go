package graphs

import (
	"fmt"
	"slices"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
)

func walkBFS(graph GraphAdjacencyMatrix, source, needle int) []int {
	visited := make([]bool, len(graph))
	parent := make([]int, len(graph)) // the parent list represent which node called the node in that position
	for i := range parent {
		parent[i] = -1
	}

	queue := algorithms.NewQueue()

	visited[source] = true
	queue.Enqueue(source)

	for queue.GetSize() != 0 {
		curr := queue.Dequeue().(int)

		if curr == needle {
			break
		}
		adjacencyRow := graph[curr]
		for i := range adjacencyRow {
			if adjacencyRow[i] == 0 {
				continue
			}

			if visited[i] {
				continue

			}

			visited[i] = true
			parent[i] = curr

			queue.Enqueue(i)
		}
	}

	if parent[needle] == -1 {
		return []int{}
	}

	fmt.Println(parent)

	curr := needle

	out := make([]int, 0)

	for parent[curr] != -1 {
		out = append(out, curr)
		curr = parent[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}

func BFSOnAdjancyMatrix(graph GraphAdjacencyMatrix, source, needle int) []int {
	return walkBFS(graph, source, needle)
}
