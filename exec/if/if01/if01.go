package main

import "fmt"

func main() {
	fmt.Printf("请输入一个数：\n")
	var num int
	fmt.Scanln(&num)
	// 判断是否是水仙花数
	if num < 100 || num > 999 {
		fmt.Printf("不是水仙花数")
	}
	a := num / 100
	b := num % 100 / 10
	c := num % 10
	fmt.Printf("a=%v,b=%v,c=%v\n", a, b, c)
	if a+b*10+c*100 == num {
		fmt.Printf("❀")
	} else {
		fmt.Printf("不是")
	}
}
