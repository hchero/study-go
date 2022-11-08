package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, 10, 20, 30)
	var arrInt = []int{1, 2, 3}
	slice = append(slice, arrInt...)
	fmt.Println(slice)
}
