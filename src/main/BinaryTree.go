package main

import (
	"../structure/queue"
	"../structure/tree"
	"fmt"
)

func generateTree(val []int) (*tree.Tree, bool) {
	fifo := queue.Constructor()
	root := tree.Construct()
	fifo.Put(root)
	var info *tree.Tree
	for i, num := range val {
		node := tree.Init(num)
		if i%2 == 0 {
			info = fifo.Get(false)
			info.AddLeft(node)
		} else {
			info = fifo.Get(true)
			info.AddRight(node)
		}
		fifo.Put(node)
	}
	isSymmetric := info.IsSymmetric()
	return info, isSymmetric
}

func main() {
	fmt.Print(generateTree([]int{1, 1, 2, 3, 3, 2}))
	fmt.Print("\n __________________ \n")
	fmt.Print(generateTree([]int{1, 1, 1, 1}))
}
