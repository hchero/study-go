package main

import "fmt"

func main() {
	// go的for循环两种写法
	fmt.Printf("第一种写法：\n")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Printf("\n第二种写法：\n")
	j := 1
	for j <= 10 {
		fmt.Printf("%d", j)
		j++
	}
	k := 1
	fmt.Printf("\n第三种写法：\n")
	for {
		if k > 10 {
			break
		}
		fmt.Printf("%d", k)
		k++
	}
}
