package main

import (
	"fmt"
)

// 使用数组
func main() {
	var (
		length            int
		n, times, starter int
	)
	fmt.Println("约瑟夫环问题\n请输入人数 报数周期 报数起始人")
	n, _ = fmt.Scanln(&length, &times, &starter)
	if length <= 0 || n == 0 {
		fmt.Println("人数有误")
		return
	} else if starter >= length || starter <= 0 {
		fmt.Println("起始位置有误")
		return
	} else if times <= 0 {
		fmt.Println("报数周期有误")
	}
	// generate people array
	var arr []int
	var result []int
	for i := 0; i < length; i++ {
		arr = append(arr, i+1)
	}
	i, pop := starter-1, 1
	fmt.Println(arr)
	for {
		if pop == times {
			result = append(result, arr[i])
			arr = append(arr[:i], arr[i+1:]...)
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
	fmt.Println(result, arr[0])
}
