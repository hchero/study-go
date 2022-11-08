package main

import "fmt"

func test(arr *[3]int) {
	(*arr)[0] = 88
}
func main() {
	var arr = [...]int{1, 2, 3}
	test(&arr)
	fmt.Println(arr)
	fmt.Printf("arr 的地址= %v \n", &arr[0])
}
