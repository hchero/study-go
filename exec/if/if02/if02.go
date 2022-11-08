package main

import "fmt"

func main() {
	fmt.Printf("请输入年份\n")
	var year int
	fmt.Scanln(&year)
	var yes bool
	if year%100 == 0 {
		if year%400 == 0 {
			yes = true
		}
	} else {
		if year%4 == 0 {
			yes = true
		}
	}
	if yes {
		fmt.Printf("闰年")
	} else {
		fmt.Printf("不是闰年")
	}
}
