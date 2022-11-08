package main

import "fmt"

func main() {
	var num int = 1
	num2 := new(int)
	fmt.Printf("num的类型%T，值%v，地址%v\n", num, num, &num)
	fmt.Printf("num2的类型%T，值%v，地址%v", num2, num2, &num2)
}
