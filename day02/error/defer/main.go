package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	var n1 int = 10
	var n2 int = 0
	var n3 int = n1 / n2
	fmt.Println(n3)
}

func main() {
	test()
	fmt.Println("结束")
}
