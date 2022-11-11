/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-11 16:36:44
 * @LastEditors: super
 * @Description:根据行、列、字符打印对应行数和列数的字符，比如：行：3，列：2，字符 *，则打印相应的效果
 */
package main

import "fmt"

type MethodUtils struct {
}

func (method *MethodUtils) print(n int, m int, key string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf(key)
		}
		fmt.Println()
	}
}

func main() {
	var method MethodUtils
	method.print(5, 8, "♥")
}
