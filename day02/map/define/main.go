package main

import "fmt"

func main() {
	// 第一种方式
	var a map[string]string
	a = make(map[string]string, 10)
	a["武松"] = "武松"
	fmt.Println(a)
	// 第二种方式
	b := make(map[string]string)
	b["no1"] = "张三"
	b["no1"] = "李四"
	b["no1"] = "王五"
	fmt.Println(b)

	// 第三种方式
	c := map[string]string{
		"no1": "张三",
		"no2": "李四",
		"no3": "王五",
	}
	fmt.Println(c)
}
