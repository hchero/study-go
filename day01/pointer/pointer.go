package main

import "fmt"

func main() {
	var i int = 10
	/*
		ptr是一个指针变量
		ptr的类型是*int
		ptr本身的值&i
	*/
	var ptr *int = &i
	// fmt.Printf("i的地址 = %v", &i)
	fmt.Printf("ptf=%v\n", ptr)
	// 取出指针指向的值
	fmt.Printf("ptf=%v\n", *ptr)
}
