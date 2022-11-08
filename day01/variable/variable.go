package main

import "fmt"

func main() {
	var i int
	// i = 10;
	// int类型的默认值是0；
	// 根据值自行判断变量类型（类型推导）
	// 省略var :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	var num = 10.11
	fmt.Println("变量i=", i)
	str := "hehe"
	fmt.Println("变量num=", num)
	fmt.Println("变量str=", str)
	// 一次性声明多个变量
	// var n1,n2,n3 int

	var (
		n1 = 1
		n2 = 2
		n3 = 3
	)

	name, page, sex := "lucy", 18, "男"
	fmt.Println(n1, n2, n3)
	fmt.Println(name, page, sex)
	var num2 int = 1
	fmt.Println(num2)

}
