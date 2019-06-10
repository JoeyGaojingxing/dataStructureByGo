package queue

import "../tree"

type queue []*tree.Tree

func Constructor() *queue {
	var res queue
	return &res
}

func (val *queue) Get(del bool) *tree.Tree {
	res := (*val)[0]
	var queue queue
	if del {
		for i := 1; i < len(*val); i++ {
			queue = append(queue, (*val)[i])
		}
		*val = queue
	}
	return res
}

func (val *queue) Put(tree *tree.Tree) {
	*val = append(*val, tree)
}
