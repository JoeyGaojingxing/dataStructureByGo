package main

import "fmt"

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
			if this.data == 0 {
				return nil
			}
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
	if val == nil {
		return -1
	} else {
		return val.data
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if val == 0 {
		return
	}
	if this.data == 0 {
		this.data = val
	} else {
		list := &MyLinkedList{data: this.data, Next: this.Next}
		this.data = val
		this.Next = list
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if val == 0 {
		return
	}
	for {
		if this.data == 0 {
			this.data = val
			break
		} else if this.Next == nil {
			tail := &MyLinkedList{data: val}
			tail.data = val
			this.Next = tail
			break
		} else {
			this = this.Next
		}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if val == 0 {
		return
	}
	pre := this
	for i := 0; ; i++ {
		if index == 0 {
			this.AddAtHead(val)
			break
		} else if i == index {
			next := &MyLinkedList{data: val, Next: this}
			pre.Next = next
			break
		} else if this != nil && this.data != 0 {
			pre = this
			this = this.Next
		} else {
			break
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	pre := this.GetPoint(index - 1)
	next := this.GetPoint(index + 1)
	if pre != nil {
		pre.Next = next
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

//func main() {
//	obj := Constructor()
//	obj.AddAtHead(0)
//	obj.AddAtIndex(1, 9)
//	obj.AddAtIndex(1, 5)
//	obj.AddAtTail(7)
//	obj.AddAtHead(1)
//	obj.AddAtIndex(5, 8)
//	obj.AddAtIndex(5, 2)
//	obj.AddAtIndex(3, 0)
//	for i:=0; i<10; i++ {
//		fmt.Print(obj.Get(i))
//	}
//	obj.AddAtTail(1)
//	obj.AddAtTail(0)
//	obj.DeleteAtIndex(6)
//	fmt.Println("end", obj.Get(3))
//}
func main() {
	obj := Constructor()
	for i := 0; i < 5; i++ {
		obj.AddAtTail(i)
		obj.AddAtHead(i)
		obj.AddAtIndex(5, i)
	}
	for i := 0; i < 10; i++ {
		fmt.Print(" |", obj.Get(i))
	}
}
