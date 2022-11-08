package main

import "fmt"

func main() {
	var a int = 1
	var b int = -1
	var c = a | b

	/*
		1的补码 	0000 0001
		-1的原码 	1000 0001
		-1的反码 	1111 1110
		-1的补码 	1111 1111

		0000 0001
		1111 1111

		补码 	1111 1111
		反码    1111 1110
		原码 	1000 0001
		-1
	*/
	fmt.Printf("%v |  %v = %v", a, b, c)
}
