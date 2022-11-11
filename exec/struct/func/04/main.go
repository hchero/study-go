/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-11 16:32:20
 * @LastEditors: super
 * @Description:编写方法：判断一个数是奇数还是偶数
 */
package main

import "fmt"

type MethodUtils struct {
}

func (method *MethodUtils) oddOrEven(n int) {
	if n%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("是奇数")

	}
}

func main() {
	var mu MethodUtils
	mu.oddOrEven(2)
}
