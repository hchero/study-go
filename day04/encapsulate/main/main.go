/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-14 07:49:17
 * @LastEditors: super
 * @Description:
 */
package main

import (
	"fmt"
	"go_code/project01/day04/encapsulate/model"
)

func main() {
	p := model.NewPerson("小明")
	p.SetAge(18)
	p.SetSal(4000)
	fmt.Println(p)
	fmt.Println(p.GetAge())
	fmt.Println(p.GetSal())
}
