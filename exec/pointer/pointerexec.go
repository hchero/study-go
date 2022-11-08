package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int
	ptr = &num
	*ptr = 9
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr=%v", *ptr)
}
