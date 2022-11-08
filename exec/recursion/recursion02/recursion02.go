package main

import "fmt"

func hsz(n int) int {
	if n == 1 {
		return 3
	} else {
		return hsz(n-1)*2 + 1
	}
}
func main() {
	fmt.Println("请输入一个数")
	var num int
	fmt.Scanln(&num)
	fmt.Println(hsz(num))
}
