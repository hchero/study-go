/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-11 16:15:17
 * @LastEditors: super
 * @Description:编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
 */

package main

import "fmt"

type MethodUtils struct {
}

func (method MethodUtils) func1(m int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	var method = MethodUtils{}
	method.func1(2, 4)
}
