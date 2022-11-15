/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-15 08:57:34
 * @LastEditors: super
 * @Description:
 */
package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p) // 输出：2
}
