/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-11 16:28:18
 * @LastEditors: super
 * @Description: 编写一个方法算该矩形的面积（可以接收长len，宽width），将其作为方法返回值。在main方法中调用该方法，接收返回的面积并打印
 */
package main

import "fmt"

type MethodUtils struct {
}

func (method MethodUtils) area(len int, width int) int {
	return len * width
}
func main() {
	var method = MethodUtils{}
	area := method.area(10, 8)
	fmt.Println("面积=", area)
}
