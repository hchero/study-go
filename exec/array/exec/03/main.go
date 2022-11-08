/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-07 21:20:39
 * @LastEditors: super
 * @LastEditTime: 2022-11-07 21:53:14
 * @Description:定义一个3行4列的二位数组，逐个从键盘输入值，编写程序将四周的数据清0
 */
package main

import "fmt"

func main() {
	var arr [3][4]int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("请输入第%v行数据：\n", i+1)
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("第%v个数：\n", j+1)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if i == 0 || i == len(arr)-1 {
			for j := 0; j < len(arr[i]); j++ {
				arr[i][j] = 0
			}
		}
		arr[i][0] = 0
		arr[i][len(arr[i])-1] = 0
	}
	fmt.Println(arr)
}
