package main

import "fmt"

func main() {
	var a int = 2
	var b int = -1
	var c int = a & b

	/*
		2的补码 	0000 0010
		-1的原码 	1000 0001
		-1的反码 	1111 1110
		-1的补码 	1111 1111

		0000 0010
		1111 1111

		补码 	0000 0010
		反码    0000 0010
		原码 	0000 0010
		2
	*/

	fmt.Printf("%v * %v = %v", a, b, c)
}
