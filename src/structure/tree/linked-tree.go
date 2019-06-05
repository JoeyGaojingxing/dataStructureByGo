package tree

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

func (tree *Tree) AddLeft(val Val) {
	add := Tree{val: val}
	tree.ltree = &add
}

func (tree *Tree) AddRight(val Val) {
	add := Tree{val: val}
	tree.rtree = &add
}

func (tree *Tree) IsSymmetric() bool {
	return isMirror(tree.ltree, tree.rtree)
}

func isMirror(ltree, rtree *Tree) bool {
	if ltree == nil && rtree == nil {
		return true
	} else if ltree == nil || rtree == nil {
		return false
	} else {
		return ltree.val == rtree.val && isMirror(ltree.ltree, rtree.rtree) && isMirror(ltree.rtree, rtree.ltree)
	}
}

func (tree *Tree) Tree2Arr() []Val {
	return []Val{}
}

func main() {

}
