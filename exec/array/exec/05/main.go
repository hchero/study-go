/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-07 22:13:44
 * @LastEditors: super
 * @LastEditTime: 2022-11-07 22:17:12
 * @Description:试保存1 3 5 7 9五个奇数到数组，并倒叙打印
 */
package main

import "fmt"

func main() {
	var arr [5]int
	for i := 1; i <= 5; i++ {
		arr[i-1] = i*2 - 1
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[len(arr)-1-i])
	}
}
