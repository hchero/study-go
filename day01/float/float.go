/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-17 14:54:16
 * @LastEditors: super
 * @Description:
 */
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num float32 = .1234
	fmt.Println(num)
	fmt.Printf("数字num的长度 %v", unsafe.Sizeof(num))
}
