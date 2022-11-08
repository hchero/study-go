package main

import "fmt"

func main() {
	var flow int
	fmt.Printf("请输入金字塔层数\n")
	fmt.Scanf("%d", &flow)
	// for i := 1; i <= flow; i++ {
	// 	// 先打印空格
	// 	for j := 1; j <= flow-i; j++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	i := 1
	for {
		if i > flow {
			break
		}
		j := 1
		for {
			if j > flow-i {
				break
			}
			fmt.Printf(" ")
			j++
		}
		k := 1
		for {
			if k > 2*i-1 {
				break
			}
			fmt.Printf("*")
			k++
		}
		i++
		fmt.Printf("\n")
	}
}
