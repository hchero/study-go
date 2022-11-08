/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-08 13:17:04
 * @LastEditors: super
 * @LastEditTime: 2022-11-08 14:51:16
 * @Description:定义一个数组，并给出8个整数，求该数组中大于平均值的个数，和小于平均值的个数
 */
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
 * @description: 求最大值，最小值，平均值，
 * @param {[]int} arr
 * @return {*}
 */
func maxAndMin(arr []int) (max int, min int, avg float64) {
	max = arr[0]
	min = arr[0]
	total := arr[0]
	for i := 1; i < len(arr); i++ {
		total += arr[i]
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	avg, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(total)/float64(len(arr))), 64)

	return max, min, avg
}

/**
 * @description: 返回小于平均数的个数，大于平均值的个数
 * @param {[]int} arr
 * @param {float64} avg
 * @return {*}
 */
func smallBigCount(arr []int, avg float64) (smallCount int, bigCount int) {
	smallCount = 0
	bigCount = 0
	for i := 0; i < len(arr); i++ {
		if float64(arr[i]) < avg {
			smallCount++
		}
		if float64(arr[i]) > avg {
			bigCount++
		}
	}
	return smallCount, bigCount
}

func main() {
	var arr []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 8; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	// 取出最大值最小值
	// 取出平均值

	max, min, avg := maxAndMin(arr)
	fmt.Printf("最大值%v,最小值%v,平均值%v\n", max, min, avg)
	// 取出小于平均值的数，取出大于平均值的数
	smallCount, bigCount := smallBigCount(arr, avg)
	fmt.Printf("小于平均值的个数%v，大于平均值的个数%v", smallCount, bigCount)
}
