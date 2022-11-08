package main

import "fmt"

func main() {
	students := make(map[string]string)
	students["no1"] = "李白"
	students["no2"] = "张三"
	value, findRes := students["no1"]
	fmt.Println(value, findRes)
	v2, f2 := students["no3"]
	fmt.Println(v2, f2)

}
