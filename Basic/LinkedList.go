package main

import "fmt"

type MyLinkedList struct {
	data int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() *MyLinkedList {
	return nil
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this == nil {
		return -1
	} else {
		for i := 0; ; i++ {
			if i == index {
				return this.data
			} else if this.Next != nil {
				this = this.Next
			} else {
				return -1
			}
		}
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this == nil {
		this = &MyLinkedList{data: val}
		return
	} else {
		node := this
		this.data = val
		this.Next = node
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this == nil {
		this = &MyLinkedList{data: val}
		return
	}
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
	if this == nil {
		if index == 0 {
			this = &MyLinkedList{data: val}
			return
		} else {
			return
		}
	} else {
		pre := this
		for i := 0; ; i++ {
			if i == index {
				if pre == this {
					this.AddAtHead(val)
					return
				} else {
					node := &MyLinkedList{data: val, Next: this}
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
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this == nil {
		return
	} else if index == 0 {
		if this == nil {
			return
		} else if this.Next == nil {
			this = nil
			return
		} else {
			this = this.Next
			return
		}
	} else {
		pre := this
		for i := 0; ; i++ {
			if i == index {
				if this.Next == nil {
					pre.Next = nil
					return
				} else {
					pre.Next = this.Next
					return
				}
			} else {
				pre = this
				this = this.Next
			}
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
