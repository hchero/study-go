/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-08 11:32:16
 * @LastEditors: super
 * @LastEditTime: 2022-11-08 12:26:55
 * @Description:随机生成10个整数（1-100），使用冒泡排序法进行排序，然后使用二分法查找，查找是否有90这个数，并显示下标，如果没有则提示“找不到该数”
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @description: 二分法查找
 * @param {[]int} arr
 * @param {int} left
 * @param {int} right
 * @param {int} findValue
 * @return {*}
 */
func dichotomyFind(arr []int, left int, right int, findValue int) int {
	var middle = (left + right) / 2
	if right < left || findValue > arr[len(arr)-1] || findValue < arr[0] {
		return -1
	}
	if arr[middle] == findValue {
		return middle
	} else if arr[middle] > findValue {
		return dichotomyFind(arr, left, middle-1, findValue)
	} else {
		return dichotomyFind(arr, middle+1, right, findValue)
	}
}

/**
 * @description: 冒泡排序
 * @param {[]int} arr
 * @return {*}
 */
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	return arr
}

func main() {
	var arr []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	// arr = []int{68, 88, 1, 47, 50, 92, 72, 4, 54, 33}
	fmt.Println("数组是：", arr)
	arr = bubbleSort(arr)
	fmt.Println("排序后：", arr)

	findIndex := dichotomyFind(arr, 0, len(arr), 99)
	if findIndex > -1 {
		fmt.Println(findIndex)
	} else {
		fmt.Println("找不到该数", findIndex)
	}
}
