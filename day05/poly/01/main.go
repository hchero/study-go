/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-16 15:26:01
 * @LastEditors: super
 * @Description:多态案例
 */
package main

import (
	"fmt"
)

/**
 * @description: 声明接口
 * @return {*}
 */
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作..")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作..")
}

type Camera struct {
	Name string
}

/**
 * @description: 实现Usb接口方法Start()
 * @return {*}
 */
func (c Camera) Start() {
	fmt.Println("相机开始工作..")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作..")
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Camera{"尼康"}
	usbArr[2] = Phone{"iphone"}

	fmt.Println(usbArr)
}
