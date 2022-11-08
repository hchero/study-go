package main

import "fmt"

func main() {
	fmt.Printf("请输入月份\n")
	var month byte
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Printf("春季")
	case 6, 7, 8:
		fmt.Printf("夏季")
	case 9, 10, 11:
		fmt.Printf("秋季")
	case 12, 1, 2:
		fmt.Printf("冬季")
	default:
		fmt.Printf("月份错误")
	}
}
