package main

import "fmt"

func test(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	// 当函数执行完成后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("defer1 n=", n1)
	defer fmt.Println("defer2 n=", n2)
	res := n1 + n2
	fmt.Println("defer3 res=", res)
	return res
}
func main() {
	res := test(1, 2)
	fmt.Println("res=", res)
}
