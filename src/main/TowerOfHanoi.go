package main

import "fmt"

func moveHanoi(num int) int {
	if num == 1 {
		return 1
	} else {
		return moveHanoi(num-1)*2 + 1
	}
}

func main() {
	var num int
	_, _ = fmt.Scanln(&num)
	result := moveHanoi(num)
	fmt.Println(result)
}
