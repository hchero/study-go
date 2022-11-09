/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-09 14:39:02
 * @LastEditors: super
 * @LastEditTime: 2022-11-09 14:43:47
 * @Description: 内存分布
 */
package main

import "fmt"

type Point struct {
	x int
	y int
}

// type Rect struct {
// 	leftUp  Point
// 	rightDown Point
// }
type Rect struct {
	leftUp, RightDown Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("内存分布查看\n")
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p\n", &r1.leftUp.x, &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n", &r1.RightDown.x, &r1.RightDown.y)
}
