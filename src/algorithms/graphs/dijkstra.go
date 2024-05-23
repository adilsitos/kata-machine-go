package graphs

import (
	"math"
	"slices"
)

func hasUnvisited(seen []bool, distances []int) bool {
	for i := range seen {
		if !seen[i] && distances[i] < math.MaxInt {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, distances []int) int {
	idx := -1
	lowestDistance := math.MaxInt

	for i := range seen {
		if seen[i] {
			continue
		}

		if lowestDistance > distances[i] {
			lowestDistance = distances[i]
			idx = i
		}
	}

	return idx
}

func DijkstraShortestPath(source, needle int, graph GraphAdjacencyList) []int {
	seen := make([]bool, len(graph))

	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	distances := make([]int, len(graph))
	for i := range distances {
		distances[i] = math.MaxInt
	}

	distances[source] = 0

	for hasUnvisited(seen, distances) {
		curr := getLowestUnvisited(seen, distances)
		seen[curr] = true

		adjs := graph[curr]
		for i := range adjs {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := distances[curr] + edge.Weight
			if dist < distances[edge.To] {
				distances[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	out := make([]int, 0)
	curr := needle

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}
