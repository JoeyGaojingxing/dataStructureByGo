package main

import (
	"C"
	"fmt"
	"unsafe"
)

//export NinePatch
func NinePatch(edge byte) unsafe.Pointer {
	fmt.Println(edge)
	//init table
	var arr [][]byte
	if edge%2 != 1 {
		fmt.Println("请输入奇数")
		return nil
	} else if edge <= 0 {
		fmt.Println("请输入大于0的数")
		return nil
	}
	var temp []byte
	// init edge * edge slice
	var i byte
	for i = 0; i < edge; i++ {
		temp = nil
		var j byte
		for j = 0; j < edge; j++ {
			temp = append(temp, 0)
		}
		arr = append(arr, temp)
	}
	now := []byte{1, (edge + 1) / 2} // 当前位置
	for i = 1; i <= edge*edge; i++ {
		arr[now[0]-1][now[1]-1] = i
		// go up-right
		now[0] = now[0] - 1
		now[1] = now[1] + 1
		fmt.Println(now)
		// judge row out of range
		if now[1] > edge {
			// if row, judge both row and col out of range
			if now[0] < 1 {
				// if both row and col, go pre-node-down
				now[0] = now[0] + 2
				now[1]--
			} else {
				// if only row, go row-start
				now[1] = 1
			}
			// judge col out of range
		} else if now[0] < 1 {
			// if col, go col-bottom
			now[0] = edge
			// if has value
		} else if arr[now[0]-1][now[1]-1] != 0 {
			// go pre-node-down
			now[0] = now[0] + 2
			now[1]--
		}
	}
	fmt.Println(arr)
	var res []byte
	for _, list := range arr {
		res = append(res, list...)
	}
	return C.CBytes(res)
}

func main() {
	res := NinePatch(3)
	fmt.Println(res)
}

func main1() {
	//init table
	var x, err int
	var arr [][]int
	fmt.Print("请输入奇数：")
	err, _ = fmt.Scan(&x)
	if x%2 != 1 {
		fmt.Println("请输入奇数")
		return
	} else if x <= 0 || err == 0 {
		fmt.Println("请输入大于0的数")
		return
	}
	var temp []int
	// init x * x slice
	for i := 0; i < x; i++ {
		temp = nil
		for j := 0; j < x; j++ {
			temp = append(temp, 0)
		}
		arr = append(arr, temp)
	}
	now := []int{0, (x - 1) / 2} // 当前位置
	for i := 1; i <= x*x; i++ {
		arr[now[0]][now[1]] = i
		// go up-right
		now[0] = now[0] - 1
		now[1] = now[1] + 1
		// judge row out of range
		if now[1] > x-1 {
			// if row, judge both row and col out of range
			if now[0] < 0 {
				// if both row and col, go pre-node-down
				now[0] = now[0] + 2
				now[1]--
			} else {
				// if only row, go row-start
				now[1] = 0
			}
			// judge col out of range
		} else if now[0] < 0 {
			// if col, go col-bottom
			now[0] = x - 1
			// if has value
		} else if arr[now[0]][now[1]] != 0 {
			// go pre-node-down
			now[0] = now[0] + 2
			now[1]--
		}
	}
	var res []int
	for _, list := range arr {
		res = append(res, list...)
	}
	fmt.Println(res)
}
