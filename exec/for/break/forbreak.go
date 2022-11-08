package main

import "fmt"

func main() {
	var money float32 = 100000
	var times int
	for {
		fmt.Printf("当前金额%v,", money)
		if money > 50000 {
			money -= money * 0.05
		} else {
			money -= 1000
		}
		if money > 0 {
			times++
			fmt.Printf("第%v次缴纳,剩余金额%v\n", times, money)
		} else {
			break
		}
	}
	fmt.Printf("经过路口%v次", times)
}
