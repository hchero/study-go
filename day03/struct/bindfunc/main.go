/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-09 18:18:30
 * @LastEditors: super
 * @LastEditTime: 2022-11-09 18:21:31
 * @Description:给结构体绑定方法
 */
package main

import "fmt"

type Person struct {
	Name string `json: "name`
	Age  int    `json: "age"`
}

func (p Person) test() {
	fmt.Println("person.Name=", p.Name)
}

func main() {
	var p Person
	p.test()
	p.Name = "super"
	p.test()
}
