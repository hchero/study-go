package main

import "fmt"

func main() {
	fmt.Println("请输入整数，0表示输入结束")
	var num int
	var positive int
	var negative int
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positive++
		} else {
			negative++
		}
		continue

	}
	fmt.Printf("正数数量%v,负数数量%v", positive, negative)
}
