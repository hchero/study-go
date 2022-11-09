/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-09 11:27:56
 * @LastEditors: super
 * @LastEditTime: 2022-11-09 11:34:22
 * @Description:结构体的定义
 */
package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64        // 数组
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // map
}

func main() {
	var p1 Person

	p1.ptr = &p1.Age
	p1.slice = make([]int, 10)
	p1.map1 = make(map[string]string)
	fmt.Println(p1)
}
