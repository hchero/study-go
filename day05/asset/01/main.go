/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-16 15:50:02
 * @LastEditors: super
 * @Description:类型断言
 */
package main

import "fmt"

type Point struct {
	y int
	x int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	// var b = a.(Point)
	b, ok := a.(Point)
	if ok {
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println(b)

	var x interface{}
	var b2 float32 = 2.1
	x = b2 // 空接口，可以接收任何类型
	fmt.Println(x)
}
