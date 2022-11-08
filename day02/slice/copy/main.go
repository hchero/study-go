package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var slice = make([]int, 1)
	fmt.Println(a)
	fmt.Println(slice)
	copy(slice, a)
	fmt.Println(slice)
}
