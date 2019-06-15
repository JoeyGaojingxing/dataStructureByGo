package main

import (
    "../structure/graph"
    "../structure/stack"
)

func main() {
    // init stack
    var stack stack.Stack
    // init graph with specific nodes
    verticals := 7
    Graph := graph.Init(verticals)

    // add edges to graph
    edges := []int{1, 2, 1, 3, 1, 4, 1, 6, 2, 6, 3, 4, 3, 5, 4, 5, 4, 6, 5, 6}
    for i, v := range edges {
        var x, y int
        if i%2 == 0 {
            x = v
        } else {
            y = v
            Graph.SetEdge(x, y)
        }
    }

    // define 4 colors
    colors := [...]int{3, 5, 7, 9}

    for i := 0; i < verticals; i++ {
        neighbors := Graph.GetEdge(i)
        for _, color := range colors {
            stack.Push([4]rune{rune(i), rune(color)})
            for j, v := range neighbors {
                if v == 1 { // 找到每个边
                    if Graph.GetValue(i, j) == color {
                        stack.Pop()
                    } else {
                        continue
                    }
                }
            }
        }
    } else {

    }

}
