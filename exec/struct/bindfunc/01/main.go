/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-10 17:15:13
 * @LastEditors: super
 * @LastEditTime: 2022-11-10 17:25:51
 * @Description: 结构体方法绑定联想
 */
package main

import "fmt"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/**
 * @description: 给Person结构体添加speak方法，输入***是一个好人
 * @return {*}
 */
func (p Person) speak() {
	fmt.Printf("%v是一个好人\n", p.Name)
}

/**
 * @description: 给Person结构体添加speak方法，计算1+...+1000的结果
 * @return {*}
 */
func (p Person) jisuan() {
	var total int
	for i := 1; i <= 1000; i++ {
		total += i
	}
	fmt.Printf("计算结果是%v\n", total)
}

func (p Person) jisuan2(n int) {
	var total int
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Printf("从1-%v的总和为%v\n", n, total)
}

func (p Person) getSum(n1 int, n2 int) {
	fmt.Println(n1, "+", n2, "=", n1+n2)
}

func main() {
	var p Person
	p.Name = "小明"
	p.speak()
	p.jisuan()
	p.jisuan2(100)
	p.getSum(3, 100)
}
