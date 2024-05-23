package algorithms_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestMazeSolver(t *testing.T) {
	maze := algorithms.Matrix{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expectedPath := []algorithms.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	end := algorithms.Point{
		X: 1,
		Y: 5,
	}

	start := algorithms.Point{
		X: 10,
		Y: 0,
	}

	res := algorithms.Solve(maze, start, end)

	assert.Equal(t, expectedPath, res)
}

func TestSimpleMaze(t *testing.T) {
	maze := algorithms.Matrix{
		"xxxxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expectedPath := []algorithms.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 9, Y: 1},
		{X: 8, Y: 1},
		{X: 7, Y: 1},
		{X: 6, Y: 1},
		{X: 5, Y: 1},
		{X: 4, Y: 1},
		{X: 3, Y: 1},
		{X: 2, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
	}

	start := algorithms.Point{
		X: 10,
		Y: 0,
	}

	end := algorithms.Point{
		X: 1,
		Y: 2,
	}

	res := algorithms.Solve(maze, start, end)

	assert.Equal(t, expectedPath, res)
}
