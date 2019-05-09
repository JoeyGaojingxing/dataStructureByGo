package main

import (
	"fmt"
)

type MyLinkedList struct {
	data int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var linkedList MyLinkedList
	return linkedList
}

func (this *MyLinkedList) GetPoint(index int) *MyLinkedList {
	for i := 0; i <= index; i++ {
		if i == index {
			// this is the value
			return this
		} else if this.Next == nil {
			// if end
			return nil
		} else {
			// Next
			this = this.Next
		}
	}
	return nil
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	val := this.GetPoint(index)
	if val == nil || val.data == 0 {
		return -1
	} else {
		return val.data
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	pre := this
	if pre.data != 0 {
		pre := MyLinkedList{
			data: this.data,
			Next: this.Next,
		}
		this.data = val
		this.Next = &pre
	} else {
		this.data = val
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	for i := 0; ; i++ {
		if this.Next == nil {
			next := MyLinkedList{data: val}
			this.Next = &next
			break
		} else {
			this = this.Next
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	var pre *MyLinkedList
	for i := 0; ; i++ {
		if index == 0 {
			this.AddAtHead(val)
			break
		} else if this.Next == nil {
			if this.data == 0 && i == 0 {
				break
			} else if i == index && pre != nil {
				insert := MyLinkedList{data: val, Next: this}
				pre.Next = &insert
				break
			} else if i == index+1 {
				this.AddAtTail(val)
				break
			} else {
				continue
			}
		} else if i != index {
			pre = this
			this = this.Next
		} else if i == index {
			insert := MyLinkedList{data: val, Next: pre.Next}
			pre.Next = &insert
			break
		} else {
			// raise wrong info
			break
		}
	}
	_ = pre
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	data := this.GetPoint(index - 1)
	if data != nil {
		next := this.GetPoint(index + 1)
		if next != nil {
			data.Next = next
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
	val, index := 23, 1
	obj := Constructor()
	param1 := obj.Get(index)
	fmt.Println("0=>", param1) // -1
	obj.AddAtHead(val - 1)
	obj.AddAtHead(val - 9)
	fmt.Println("1=>", obj.Get(0)) // 23
	obj.AddAtTail(val + 1)
	fmt.Println("2=>", obj.Get(1))
	fmt.Println("result=>", obj.Get(0), obj.Get(1))
	obj.AddAtIndex(1, val)
	fmt.Println("3=>", obj.Get(index))
	fmt.Println("result=>", obj.Get(0), obj.Get(1), obj.Get(2))
	//obj.DeleteAtIndex(index)
	//fmt.Println("4=>", obj.Get(index))
	fmt.Println("result=>", obj.Get(0), obj.Get(1), obj.Get(2), obj.Get(3))
}
