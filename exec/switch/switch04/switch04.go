package main

import "fmt"

func main() {
	fmt.Printf("请输入年\n")
	var year int
	fmt.Scanln(&year)
	fmt.Printf("请输入月\n")
	var month int
	fmt.Scanln(&month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(31)
	case 2:
		// 判断是否是闰年
		if isRun(year) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	}

}

func isRun(year int) bool {
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
	return yes
}
