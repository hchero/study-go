/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-11 16:05:46
 * @LastEditors: super
 * @LastEditTime: 2022-11-11 16:26:34
 * @Description:编写结构体（MethodUtils),编写一个方法，方法不需要参数，再方法中大于一个10*8的举行，在main方法中调用该方法
 */
package main

import "fmt"

type MethodUtils struct {
}

func (m MethodUtils) rectangle() {
	width := 8
	height := 10
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var m = MethodUtils{}
	m.rectangle()
}
