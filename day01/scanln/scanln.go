package main

import "fmt"

func main() {
	var name string
	// fmt.Printf("请输入你的名字：\n")

	// fmt.Scanln(&name)
	// fmt.Printf("name = %v", name)
	var age byte
	var wages byte
	var isThrough bool

	fmt.Printf("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开\n")
	fmt.Scanf("%v %d %d %t", &name, &age, &wages, &isThrough)
	fmt.Printf("姓名=%v 年龄=%v 薪水=%v 是否通过考试=%v", name, age, wages, isThrough)

}
