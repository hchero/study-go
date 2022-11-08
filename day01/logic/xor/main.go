package main

import "fmt"

func main() {
	var a int = 1
	var b int = -1
	var c = a ^ b
	/*
		1的补码 	0000 0001
		-1的补码	1111 1111
		按位异或	1111 1110
		补码		1111 1110
		反码		1111 1101
		原码		1000 0010
		结果是		-2
	*/
	fmt.Printf("%v ^  %v = %v", a, b, c)
}
