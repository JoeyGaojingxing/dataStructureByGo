/*
约瑟夫环问题
思路：
一： 数学方式，取模计算结果。代码少，运行效率低？python低……应该和内部实现有关
二： 生成一个1~n的数组，将剔除的对象打印（或存储）并赋值为0，迭代+1前进行判断，跳过0值
三： 构建链式结构？猜测：构建结构体{内存地址：数据值}，单独列表储存内存地址序列（不知道如何实现）
Q1: 时间复杂度和空间复杂度是啥
Q2: 效率高 代码量 关联 Python和Go不同
Q3: 模拟 递归 通顶
*/
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
	for i := 2; i <= length; i++ {
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
			linkList = linkList.DeleteAtIndex(times)
			fmt.Println(linkList.Data)
		}
	}
}
