package main

import "fmt"

func main() {
	var arr [][]int
	arr = append(arr, [][]int{{1, 2, 4}, {4, 4, 5}}...)
	fmt.Println(arr)

	arr2 := make([][]int, 10)
	fmt.Println(arr2)
}
