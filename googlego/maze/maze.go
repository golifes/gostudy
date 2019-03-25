package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			//maze at next is 0
			// and steps at next is 0
			// and next != start
			if value, ok := next.at(maze); !ok || value == 1 {
				continue
			}
			if value, ok := next.at(steps); !ok || value != 0 {
				continue
			}

			if next == start {
				continue
			}
			curStems, _ := cur.at(steps)
			steps[next.i][next.j] = curStems + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("googlego/maze/maze.in")
	fmt.Println(maze)
	ret := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range ret {
		for _, val := range row {
			fmt.Printf("%3d  ", val)
		}

		fmt.Println()
	}
}
