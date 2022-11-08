package main

import "fmt"

func main() {
	var arr [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v\t", i, j, v2)
		}
		fmt.Println()
	}
}
