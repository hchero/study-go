package main

import "fmt"

func main() {
	var n1 int32 = 12
	var n2 int64
	var n3 int8
	// n2 = n1 +20 // int32--->int64 错误
	// n3 = n1 + 20 // int32 --->int8 错误
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Printf("n2=%d,n3=%d", n2, n3)
}
