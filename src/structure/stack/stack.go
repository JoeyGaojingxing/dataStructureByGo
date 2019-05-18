package stack

type stack struct {
	stack [][]int
	limit int
}

func NewStack(limit int) *stack {
	data := stack{limit: limit}
	return &data
}

func (parent stack) Push(val []int) {
	if parent.IsFull() {
		return
	}
	arr := append(parent.stack, val)
	parent.stack = arr
	return
}

func (parent stack) Pop() []int {
	if parent.IsNone() {
		return nil
	}
	length := len(parent.stack)
	result := parent.stack[length-1]
	parent.stack = parent.stack[:length-1]
	return result
}

func (parent stack) IsNone() bool {
	if len(parent.stack) == 0 {
		return true
	} else {
		return false
	}
}

func (parent stack) IsFull() bool {
	if len(parent.stack) == parent.limit {
		return true
	} else {
		return false
	}
}

func (parent stack) Peek() []int {
	if parent.IsNone() {
		return nil
	} else {
		return parent.stack[len(parent.stack)-1]
	}
}
