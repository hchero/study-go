package main

import "fmt"

func main() {
	var n1 int32 = 12
	// var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 // 编译通过，溢出，结果是32
	// n3 = int8(n1) + 128 // 编译不通过
	fmt.Println(n4)
}
