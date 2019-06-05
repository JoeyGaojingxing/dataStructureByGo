package graph

type Graph [][]int

func Init(num int) *Graph {
	var res Graph
	for i := 0; i < num; i++ {
		var mid []int
		for j := 0; j < num; j++ {
			mid = append(mid, 0)
		}
		res = append(res, mid)
	}
	return &res
}

func (val *Graph) MatchBoth(x, y int) {
	(*val)[x][y] = 1
	(*val)[y][x] = 1
}

func main() {}
