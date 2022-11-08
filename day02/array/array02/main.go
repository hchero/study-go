package main

import "fmt"

func main() {
	// var intArr [3]int
	var intArr [3]int8
	var count int = len(intArr)
	for i := 0; i < count; i++ {
		fmt.Println(intArr[i], ",地址=", &intArr[i])
	}
}
