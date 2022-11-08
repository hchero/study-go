package main

import "fmt"

/*
1,1,2,3,5,8,13......
*/
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	fmt.Println("请给一个整数")
	var num int
	fmt.Scanln(&num)
	fmt.Println(fbn(num))
}
