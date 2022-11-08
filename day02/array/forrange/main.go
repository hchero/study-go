package main

import "fmt"

func main() {
	var heroes = [...]string{"张三", "李四", "王五"}
	fmt.Println(heroes)

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
	}
	for _, v := range heroes {
		fmt.Printf("v=%v\n", v)
	}
}

// 引用 reference
