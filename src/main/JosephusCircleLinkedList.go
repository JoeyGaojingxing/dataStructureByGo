package main

import (
	structure "../structure/none-head-linked-list"
	"fmt"
)

// 使用链表
func main() {
	var (
		length            int
		n, times, starter int
	)
	fmt.Println("约瑟夫环问题\n请输入人数 报数周期 报数起始人")
	n, _ = fmt.Scanln(&length, &times, &starter)
	if length <= 0 || n == 0 {
		fmt.Println("数组长度有误")
		return
	} else if length > 1000 {
		fmt.Println("最大长度为1000")
		return
	} else if starter >= length || starter <= 0 {
		fmt.Println("起始位置有误")
		return
	} else if times <= 0 {
		fmt.Println("报数周期有误")
	}
	// generate people array
	//linkList := structure.Constructor()
	linkList := &structure.MyLinkedList{Data: 1}
	// init linked list
	var i byte
	for i = 2; int(i) <= length; i++ {
		linkList.AddAtTail(i)
	}
	// closing the circle
	linkList.Circle()
	linkList.Go(byte(starter) - 1)
	// main function
	for n := 1; ; n++ {
		if linkList == linkList.Next {
			return
		} else {
			linkList = linkList.DeleteAtIndex(byte(times))
			fmt.Println(linkList.Data)
		}
	}
}
