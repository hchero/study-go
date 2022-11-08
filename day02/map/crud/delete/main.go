package main

import "fmt"

func main() {
	student := make(map[string]string)
	student["01"] = "小明"
	student["02"] = "小白"
	student["03"] = "小花"
	fmt.Println(student)
	delete(student, "01")
	fmt.Println(student)

}
