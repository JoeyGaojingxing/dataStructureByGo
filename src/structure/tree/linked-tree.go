package tree

import "fmt"

type Tree struct {
	ltree *Tree
	val   Val
	rtree *Tree
}

type Val int

func Construct() *Tree {
	var res Tree
	return &res
}

func Init(val int) *Tree {
	return &Tree{val: Val(val)}
}

func (tree *Tree) AddLeft(val *Tree) {
	tree.ltree = val
}

func (tree *Tree) AddRight(val *Tree) {
	tree.rtree = val
}

func (tree *Tree) IsSymmetric() bool {
	return isMirror(tree.ltree, tree.rtree)
}

func isMirror(ltree, rtree *Tree) bool {
	if ltree == nil && rtree == nil {
		fmt.Print("nil, true\n")
		return true
	} else if ltree == nil || rtree == nil {
		fmt.Print("nil, false\n")
		return false
	} else {
		fmt.Print("ltree rtree :", ltree.val, rtree.val, "\n")
		return ltree.val == rtree.val && isMirror(ltree.ltree, rtree.rtree) && isMirror(ltree.rtree, rtree.ltree)
	}
}

func (tree *Tree) Tree2Arr() []Val {
	return []Val{}
}
