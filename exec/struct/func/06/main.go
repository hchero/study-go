/*
* @Author: super zhou4xin3xin3520@163.com
* @Date: 2022-11-11 16:41:13
  - @LastEditors: super

* @Description: 定义小小计算器结构体（Calculator)，实现加减乘除四个功能
实现方法1：分四个方法完成
实现方法2：用一个方法搞定
*/
package main

import "fmt"

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (c Calculator) compute(operator byte) {
	var result float64
	switch operator {
	case '+':
		result = c.Num1 + c.Num2
	case '-':
		result = c.Num1 - c.Num2
	case '*':
		result = c.Num1 * c.Num2
	case '/':
		result = c.Num1 / c.Num2
	default:
		fmt.Println("方法错误")
		return
	}
	fmt.Printf("%v %c %v = %v\n", c.Num1, operator, c.Num2, result)
}

func main() {
	var c Calculator = Calculator{1, 2}
	c.compute('+')
	c.compute('-')
	c.compute('*')
	c.compute('/')
}
