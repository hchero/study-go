/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-07 21:47:35
 * @LastEditors: super
 * @LastEditTime: 2022-11-07 22:12:35
 * @Description:定义一个4行4列的二维数组，逐个从键盘输入值，然后将第1行和第4行的数据进行交换，第2行和第3行的数据进行交换
 */
package main

import "fmt"

func main() {
	var arr [4][4]int
	for i := 0; i < len(arr); i++ {
		// fmt.Printf("请输入第%v行的数据，空格间隔", i+1)

		fmt.Printf("请输入第%v行的数据\n", i+1)
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%v个的数据", j+1)
			fmt.Scanln(&arr[i][j])
		}

	}
	fmt.Println(arr)

	tempArr := arr[0]
	arr[0] = arr[3]
	arr[3] = tempArr
	tempArr = arr[1]
	arr[1] = arr[2]
	arr[2] = tempArr
	fmt.Println(arr)
}
