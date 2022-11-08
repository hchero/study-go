package main

import "fmt"

func main() {
	var intArr = [...]int{1, 2, 3, 4, 5, 6, 7}
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println("slice的容量 = ", cap(slice))
}
