package main

import (
	"../structure/stack"
	"fmt"
	"math"
)

/*
heuristic = [[9, 8, 7, 6, 5, 4],
             [8, 7, 6, 5, 4, 3],
             [7, 6, 5, 4, 3, 2],
             [6, 5, 4, 3, 2, 1],
             [5, 4, 3, 2, 1, 0]]
*/

func generateHeuristic(grid [][]rune, goal []rune) [][]rune {
	var heuristic [][]rune
	heuristic = init2DArr(rune(len(grid)), rune(len(grid[0])))
	var i, j rune
	for i = 0; i < rune(len(grid)); i++ {
		for j = 0; j < rune(len(grid[0])); j++ {
			heuristic[i][j] = rune(math.Abs(float64(i-goal[0]))) + rune(math.Abs(float64(j-goal[1])))
			if grid[i][j] == 1 {
				heuristic[i][j] = -1
			}
		}
	}
	return heuristic
}

func init2DArr(row, col rune) [][]rune {
	var arr [][]rune
	var middle []rune
	var i, j rune
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
func AStar(grid [][]rune, init, goal []rune, cost rune) [][2]rune {
	delta := [][]rune{
		{-1, 0}, // go up
		{0, -1}, // go left
		{1, 0},  // go down
		{0, 1},  // go right
	}

	heuristic := generateHeuristic(grid, goal)
	closed := init2DArr(rune(len(grid)), rune(len(grid[0])))
	action := init2DArr(rune(len(grid)), rune(len(grid[0])))
	closed[init[0]][init[1]] = 1

	x := init[0]
	y := init[1]
	var g rune = 0
	f := g + heuristic[init[0]][init[1]]
	var cell stack.Stack
	cell.Push([...]rune{f, g, x, y})

	found := false
	resign := false

	for {
		if found || resign {
			break
		} else {
			if cell.IsEmpty() == true {
				resign = true
				return nil
			}
			next, err := cell.Pop()
			if err == false {
				break
			}
			x := next[2]
			y := next[3]
			g := next[1]
			f := next[0]
			fmt.Println("x y g f", x, y, g, f)

			if x == goal[0] && y == goal[1] {
				found = true
			} else {
				var i rune
				for i = 0; i < rune(len(delta)); i++ {
					x2 := x + delta[i][0]
					y2 := y + delta[i][1]
					if 0 <= x2 && x2 < rune(len(grid)) && 0 <= y2 && y2 < rune(len(grid[0])) {
						fmt.Println()
						if closed[x2][y2] == 0 && grid[x2][y2] == 0 {
							g2 := g + cost
							f2 := g2 + heuristic[x2][y2]
							cell.Push([4]rune{f2, g2, x2, y2})
							closed[x2][y2] = 1
							action[x2][y2] = i
						}
					}
				}
			}
		}
	}
	var invpath [][2]rune
	x = goal[0]
	y = goal[1]
	invpath = append(invpath, [...]rune{x, y})
	for {
		if x != init[0] || y != init[1] {
			x2 := x - delta[action[x][y]][0]
			y2 := y - delta[action[x][y]][1]
			x = x2
			y = y2
			invpath = append(invpath, [2]rune{x, y})
		} else {
			break
		}
	}

	var path [][2]rune
	var i rune
	for i = 0; i < rune(len(invpath)); i++ {
		path = append(path, invpath[rune(len(invpath))-1-i])
	}
	fmt.Println("Action Map")
	for i = 0; i < rune(len(action)); i++ {
		fmt.Println(action[i])
	}
	fmt.Println("f:", f)
	return path
}

func main() {
	// 0 are free path whereas 1's are obstacles
	grid := [][]rune{{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0}}
	init := []rune{0, 0}
	goal := []rune{rune(len(grid) - 1), rune(len(grid[0]) - 1)}
	cost := rune(1)
	res := AStar(grid, init, goal, cost)
	fmt.Println(res)
	for i := range res {
		fmt.Println(i)
	}
}
