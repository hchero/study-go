/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-14 07:43:23
 * @LastEditors: super
 * @Description:
 */
package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 100 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水不合法..")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
