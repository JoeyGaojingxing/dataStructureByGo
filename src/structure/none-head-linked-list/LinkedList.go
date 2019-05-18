package none_head_linked_list

type MyLinkedList struct {
	Data byte
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() *MyLinkedList {
	return nil
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index byte) byte {
	if this == nil {
		return 0
	} else {
		var i byte = 0
		for ; ; i++ {
			if i == index {
				return this.Data
			} else if this.Next != nil {
				this = this.Next
			} else {
				return 0
			}
		}
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val byte) {
	if this == nil {
		this = &MyLinkedList{Data: val}
		return
	} else {
		node := this
		this.Data = val
		this.Next = node
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val byte) {
	if this == nil {
		this = &MyLinkedList{Data: val}
		return
	}
	for {
		if this.Next == nil {
			node := &MyLinkedList{Data: val}
			this.Next = node
			return
		} else {
			this = this.Next
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index, val byte) {
	if this == nil {
		if index == 0 {
			this = &MyLinkedList{Data: val}
			return
		} else {
			return
		}
	} else {
		pre := this
		var i byte = 0
		for ; ; i++ {
			if i == index {
				if pre == this {
					this.AddAtHead(val)
					return
				} else {
					node := &MyLinkedList{Data: val, Next: this}
					pre.Next = node
					return
				}
			} else if this.Next == nil {
				if i+1 == index {
					this.AddAtTail(val)
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
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index byte) *MyLinkedList {
	if this == nil {
		return nil
	} else if index == 0 {
		if this == nil {
			return nil
		} else if this.Next == nil {
			this = nil
			return nil
		} else {
			result := this
			this = this.Next
			return result
		}
	} else {
		pre := this
		var i byte = 1
		for ; ; i++ {
			if i == index {
				if this.Next == nil {
					result := this
					pre.Next = nil
					return result
				} else if this == this.Next {
					result := this
					this = nil
					return result
				} else {
					pre.Next = this.Next
					return this
				}
			} else {
				pre = this
				this = this.Next
			}
		}
	}
}

// close the linked list to circle
func (this *MyLinkedList) Circle() {
	start := this
	for {
		if this.Next == nil {
			this.Next = start
			return
		} else {
			this = this.Next
		}
	}
}

func (this *MyLinkedList) Go(val byte) *MyLinkedList {
	var i byte = 1
	for ; i < val; i++ {
		this = this.Next
	}
	return this
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
	val, index := 5, 0
	obj := Constructor()
	fmt.Printf("%T", obj)
	obj = Constructor()
	param_1 := obj.Get(index)
	obj.AddAtHead(val)
	obj.AddAtTail(val)
	obj.AddAtIndex(index, val)
	obj.DeleteAtIndex(index)
	fmt.Println(param_1)
}
*/
