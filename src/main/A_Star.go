package main

import (
	"math"
	"../structure/stack"
)

// 0 are free path whereas 1's are obstacles
grid = [[0, 1, 0, 0, 0],
[0, 1, 0, 0, 0],
[0, 1, 0, 0, 0],
[0, 1, 0, 1, 0],
[0, 0, 0, 1, 0]]

/*
heuristic = [[9, 8, 7, 6, 5, 4],
             [8, 7, 6, 5, 4, 3],
             [7, 6, 5, 4, 3, 2],
             [6, 5, 4, 3, 2, 1],
             [5, 4, 3, 2, 1, 0]]
*/

func generateHeuristic(grid [][]byte, goal []byte) [][]byte {
	var heuristic [][]byte
	heuristic = init2DArr(byte(len(grid)), byte(len(grid[0])))
	var i, j byte
	for i = 0; i < byte(len(grid)); i++ {
		for j = 0; j < byte(len(grid[0])); j++ {
			heuristic[i][j] = byte(math.Abs(float64(i-goal[0]))) + byte(math.Abs(float64(j-goal[1])))
			if grid[i][j] == 1 {
				heuristic[i][j] = -1
			}
		}
	}
	return heuristic
}

func init2DArr(row, col byte) [][]byte {
	var arr [][]byte
	var middle []byte
	var i, j byte
	for i = 0; i < row; i++ {
		middle = nil
		for j = 0; j < col; j++ {
			middle = append(middle, 0)
		}
		arr = append(arr, middle)
	}
	return arr
}

//export AStar
func AStar(grid [][]byte, init, goal []byte, cost byte) {
	heuristic := generateHeuristic(grid, goal)
	closed := init2DArr(byte(len(grid)), byte(len(grid[0])))
	action := init2DArr(byte(len(grid)), byte(len(grid[0])))
	closed[init[0]][init[1]] = 1

	x := init[0]
	y := init[1]
	var g byte = 0
	f := g + heuristic[init[0]][init[1]]

	found := false
	resign := false

	for {
		if found && resign {
			break
		} else {
			next :=
		}
	}
	return
}
