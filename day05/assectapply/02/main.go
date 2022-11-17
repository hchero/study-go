/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-17 15:52:41
 * @LastEditors: super
 * @Description:断言的最佳实践2
 */
package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是 float32类型，值是%v\n", index, x)
		case int, int16, int8, int32, int64:
			fmt.Printf("第%v个参数是 int类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 float32类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是 类型不确定，值是%v\n", index, x)
		}

	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int64 = 30
	var n4 string = "super"
	var n5 int = 500
	var n6 bool = true
	TypeJudge(n1, n2, n3, n4, n5, n6)
}
