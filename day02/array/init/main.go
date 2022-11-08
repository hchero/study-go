package main

import "fmt"

func main() {
	var numArr01 [3]int = [3]int{1, 2, 3}
	var numArr02 = [3]int{1, 2, 3}
	var numArr03 = [...]int{1, 2, 3}
	var numArr04 = [...]int{1: 88, 0: 9, 2: 3}
	fmt.Println(numArr01)
	fmt.Println(numArr02)
	fmt.Println(numArr03)
	fmt.Println(numArr04)
}
