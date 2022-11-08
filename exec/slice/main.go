package main

import "fmt"

func fbn(n int) []int {
	var arr []int
	arr = append(arr, 1, 1)
	for i := 2; i < n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
		// arr = append(arr, dg(i))
	}
	return arr
}

func dg(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return dg(n-1) + dg(n-2)
}

func main() {
	fmt.Println("请输入一个数字")
	var n int
	fmt.Scanln(&n)
	arr := fbn(n)
	fmt.Println(arr)
}
