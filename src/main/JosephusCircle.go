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
	"fmt"
)

// 使用数组
func main() {
	var (
		length   int
		n, times int
	)
	fmt.Println("约瑟夫环问题\n请输入人数 报数周期")
	n, _ = fmt.Scanln(&length, &times)
	if length == 0 || n == 0 {
		fmt.Println("数组长度有误")
		return
	} else if length > 1000 {
		fmt.Println("最大长度为1000")
		return
	}
	// generate people array
	var arr []int
	var result []int
	for i := 0; i < length; i++ {
		arr = append(arr, i+1)
	}
	i, pop := 0, 1
	fmt.Println(len(arr))
	for {
		if pop == times {
			result = append(result, arr[i])
			fmt.Println(arr[i])
			arr = append(arr[0:i], arr[i+1:]...)
			pop = 0
			i--
		}
		if len(result) == length-1 {
			fmt.Println("最后一人", arr[0])
			break
		} else if len(arr) <= i+1 {
			i = 0
		} else {
			i++
		}
		pop++
	}
	fmt.Println(result)
}

// 使用链表
/*
func main() {

}
*/
