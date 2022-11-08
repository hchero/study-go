package main

import "fmt"

func PrintPyramid(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func main() {
	PrintPyramid(5)
}
// 乘法表
// multiplication