package head_linked_list

type MyLinkedList struct {
	data int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() *MyLinkedList {
	node := MyLinkedList{}
	return &node
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.Next != nil {
		this = this.Next
	} else {
		return -1
	}
	for i := 0; ; i++ {
		if i == index {
			return this.data
		} else if this.Next == nil {
			return -1
		} else {
			this = this.Next
		}
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	next := this.Next
	node := &MyLinkedList{data: val, Next: next}
	this.Next = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	for {
		if this.Next == nil {
			node := &MyLinkedList{data: val}
			this.Next = node
			return
		} else {
			this = this.Next
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// TODO: why the value -1 of index is valid
	//if index == -1 {
	//    node := &MyLinkedList{data:val}
	//    this.Next = node
	//    return
	//}
	pre := this
	for i := 0; ; i++ {
		if i == index+1 && index >= 0 {
			node := &MyLinkedList{data: val, Next: this}
			pre.Next = node
			return
		} else if this.Next == nil {
			if i == index {
				node := &MyLinkedList{data: val}
				this.Next = node
				return
			} else {
				return
			}
		} else {
			pre = this
			this = this.Next
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	pre := this
	for i := 0; ; i++ {
		if i == index+1 && index >= 0 {
			pre.Next = this.Next
			return
		} else if this.Next == nil {
			return
		} else {
			pre = this
			this = this.Next
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/*
func main() {
	obj := Constructor()
	obj.AddAtIndex(-1, 0)
	obj.Get(0)
	obj.DeleteAtIndex(-1)
	fmt.Println(obj.Get(0))
}
*/
