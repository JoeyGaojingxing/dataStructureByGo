package graph

type graph struct {
	edge     [][]int
	vertical [][]int
}

func Init(num int) *graph {
	var res [][]int
	for i := 0; i < num; i++ {
		var mid []int
		for j := 0; j < num; j++ {
			mid = append(mid, 0)
		}
		res = append(res, mid)
	}
	return &graph{edge: res, vertical: res}
}

func (val *graph) SetEdge(x, y int) {
	val.edge[x][y] = 1
	val.edge[y][x] = 1
}

func (val *graph) GetEdge(x int) []int {
	return val.edge[x]
}

func (val *graph) SetValue(x, y, z int) {
	val.vertical[x][y] = z
}

func (val *graph) GetValue(x, y int) int {
	return val.vertical[x][y]
}
