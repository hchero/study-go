package main

import "fmt"

func main() {
	// var arr [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	arr := [][]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, arr[i][j])
		}
		fmt.Println()
	}

}
