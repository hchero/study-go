/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-08 13:09:40
 * @LastEditors: super
 * @LastEditTime: 2022-11-08 13:16:16
 * @Description:编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的数和最小的数对应的下标是多少
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @description: 找出最大最小数的下标
 * @param {[]int} arr
 * @return {*}
 */
func minAndMax(arr [5]int) (maxIndex int, minIndex int) {
	var max = arr[0]
	var min = arr[0]
	for i := 1; i <= len(arr)/2; i++ {
		if arr[i] > max {
			max = arr[0]
			maxIndex = i
		}
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	return maxIndex, minIndex
}

func main() {
	var arr [5]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	maxIndex, minIndex := minAndMax(arr)
	fmt.Printf("最大值的下标是：%v，最小值的下标是：%v", maxIndex, minIndex)
}
