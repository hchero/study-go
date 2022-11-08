package main

import "fmt"

func main() {
	var a int = 10
	var b int = 6
	/*
		补码=反码+1
		a的	补码00001010
		b的	补码00000110

		a&b =00000010 2
		a|b = 00001110 14
		a^b = 00001100 12

		-2 的原码 1000 0010
			反码  11111101
			补码  11111110

		2 的补码 0000 0010

		-2^2=	补码 11111100
				反码 11111011
				原码 10000100
				-4

	*/

	//
	var c = a & b
	var d = a | b
	var e = a ^ b
	fmt.Printf("a=%v b=%v c=%v d=%v e=%v", a, b, c, d, e)
}
