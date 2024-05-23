package algorithms

// problem statement
// Imagine you have to solve a maze
// you have a starting point and an endpoint

/// Example:

// [["######### #"],
//	["#         #"],
//  ["# #########"]]

type Point struct {
	X int // columns
	Y int // lines
}

type Matrix []string

// first things to do:

// define the base case
//	- the starting point is equal to the endpoint
//  - are the current point inside the maze
//  - check if the current point was already visited
// define the recursive case
// 	- recurse into top, right, bottom and left

var direction = [][]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

var path []Point

func FindEnd(maze Matrix, visited [][]bool, current, end Point) bool {
	// this base

	if current.X == end.X && current.Y == end.Y {
		path = append(path, current)
		return true
	}

	// check if we are inside the maze
	if current.X >= len(maze[0]) || current.X < 0 || current.Y >= len(maze) || current.Y < 0 {
		return false
	}

	// check if the current point is not a wall
	if maze[current.Y][current.X] == 'x' || maze[current.Y][current.X] == 'X' {
		return false
	}

	// check if the current point was already visited
	if visited[current.Y][current.X] {
		return false
	}

	// 1. pre recursion
	visited[current.Y][current.X] = true
	path = append(path, current)

	// recursion

	for _, dir := range direction {
		point := Point{
			X: current.X + dir[1],
			Y: current.Y + dir[0],
		}

		if FindEnd(maze, visited, point, end) {
			return true
		}
	}

	n := len(path) - 1
	path = path[:n]
	return false
}

func Solve(maze Matrix, start, end Point) []Point {
	visited := make([][]bool, len(maze))

	for i := 0; i < len(maze); i++ {
		visited[i] = make([]bool, len(maze[0]))
	}

	if len(maze) != len(visited) || len(maze[0]) != len(visited[0]) {
		panic("matrixes have differente size")
	}

	if FindEnd(maze, visited, start, end) {
		return path
	}

	return path
}
