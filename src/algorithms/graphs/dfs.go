package graphs

var path []int

func walk(graph GraphAdjacencyList, curr, needle int, visited []bool) bool {
	if visited[curr] {
		return false
	}

	// pre
	visited[curr] = true
	path = append(path, curr)

	if curr == needle {
		return true
	}

	//recursion

	for _, val := range graph[curr] {
		if walk(graph, val.To, needle, visited) {
			return true
		}
	}

	// post
	// this pop just happens if we never found the value
	path = path[:len(path)-1]
	return false

}

func DFSAdjancyList(graph GraphAdjacencyList, source, needle int) []int {
	visited := make([]bool, len(graph))

	walk(graph, source, needle, visited)
	return path
}
