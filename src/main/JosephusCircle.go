package main

import (
	structure "../structure/none-head-linked-list"
	"C"
	"fmt"
	"unsafe"
)

//export JosephusCircleArr
func JosephusCircleArr(length, times, starter byte) unsafe.Pointer {
	fmt.Println(length, times, starter)
	if length <= 0 {
		fmt.Print(length, "人数有误")
		return nil
	} else if starter >= length || starter <= 0 {
		fmt.Println(starter, "起始位置有误")
		return nil
	} else if times <= 0 {
		fmt.Println(times, "报数周期有误")
	}
	// generate people array
	var arr []byte
	var result []byte
	var i byte = 0
	for ; i < length; i++ {
		arr = append(arr, i+1)
	}
	i = starter - 1
	var pop byte = 1
	fmt.Println(arr)
	for {
		if pop == times {
			result = append(result, arr[i])
			arr = append(arr[:i], arr[i+1:]...)
			pop = 0
			i--
		}
		if byte(len(result)) == length-1 {
			break
		} else if byte(len(arr)) <= i+1 {
			i = 0
		} else {
			i++
		}
		pop++
	}
	result = append(result, arr[0])
	fmt.Println("end1", result)
	res := C.CBytes(result)
	return res
}

//export JosephusCircleLinkedList
func JosephusCircleLinkedList(length, times, starter byte) unsafe.Pointer {
	if length <= 0 {
		fmt.Println("数组长度有误")
		return nil
	} else if starter >= length || starter <= 0 {
		fmt.Println("起始位置有误")
		return nil
	} else if times <= 0 {
		fmt.Println("报数周期有误")
	}
	// generate people array
	//linkList := structure.Constructor()
	linkList := &structure.MyLinkedList{Data: 1}
	// init linked list
	var i byte = 2
	for ; i <= length; i++ {
		linkList.AddAtTail(i)
	}
	// closing the circle
	linkList.Circle()
	linkList = linkList.Go(starter)
	// main function
	var result []byte
	for {
		if linkList == linkList.Next {
			result = append(result, linkList.Data)
			break
		} else {
			linkList = linkList.DeleteAtIndex(times)
			result = append(result, linkList.Data)
			linkList = linkList.Next
		}
	}
	fmt.Println("end2", result)
	res := C.CBytes(result)
	return res
}

func main() {}
