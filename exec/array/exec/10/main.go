/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-08 14:52:19
 * @LastEditors: super
 * @LastEditTime: 2022-11-09 11:18:21
 * @Description:跳水比赛，8个评委打分。运动员的成绩是8个成绩去掉一个最高分，去掉一个最低分，剩下的6个分数的平均分就是最后得分。
 */
/*
使用一维数组实现如下功能
 1、请把打最高分的评委和最低分的评委找出来。
 2、找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委，最差评委就是打分和最后得分相差最大的。
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func avgScore(arr []float64) float64 {
	max := arr[0]
	min := arr[0]
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
	total = total - max - min
	avg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", total/float64(6)), 64)
	return avg
}

/**
 * @description: 递归排序
 * @param {[]float64} arr
 * @return {*}
 */
func dichotomySort(arr []float64) []float64 {
	if len(arr) < 2 {
		return arr
	}
	first := arr[0]
	var left []float64
	var right []float64
	for i := 0; i < len(arr); i++ {
		if first > arr[i] {
			left = append(left, arr[i])

		}
		if first < arr[i] {
			right = append(right, arr[i])
		}
	}
	var sortArr []float64
	left = dichotomySort(left)
	right = dichotomySort(right)
	if len(left) >= 1 {
		sortArr = append(left, first)
	} else {
		sortArr = append(sortArr, first)
	}
	if len(right) >= 1 {
		sortArr = append(sortArr, right...)
	}

	return sortArr
}

func dichotomyIndex(arr []float64, findVal float64) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		fmt.Println("start=", start, "end", end)
		middle := (start + end) / 2
		if findVal == arr[middle] {
			return middle
		} else if findVal > arr[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return start

}

func main() {
	var arr []float64
	fmt.Println("请评委老师开始打分")
	for i := 0; i < 8; i++ {
		arr = append(arr, 0)
		fmt.Printf("请第%v位评论打分：\n", i+1)
		fmt.Scanln(&arr[i])
	}
	// arr := []float64{89, 93, 95, 87, 100, 80, 94, 93}
	// arr := []float64{89, 93, 95, 87, 100, 80, 94, 92}
	// 去掉最高分，去掉最低分
	// 求平均分
	fmt.Println("打分结果：", arr)
	avg := avgScore(arr)
	fmt.Println("平均分是", avg)
	// 递归排序
	sortArr := dichotomySort(arr)
	fmt.Println("排序后：", sortArr)
	// 查找最近的
	// 最接近 closest
	closest := dichotomyIndex(sortArr, avg)
	fmt.Println("最接近的是：", closest)
	// 最好的 最差的
	best := closest
	poorStart := math.Abs(sortArr[closest-1] - avg)
	poorEnd := math.Abs(sortArr[closest] - avg)
	if poorStart > poorEnd {
		best = closest
		fmt.Printf("最佳评委是：评分%v的评委\n", sortArr[best])
	} else if poorStart < poorEnd {
		best = closest - 1
		fmt.Printf("最佳评委是：评分%v的评委\n", sortArr[best])
	} else {
		fmt.Printf("最佳评委是：评分%v、%v的评委\n", sortArr[closest], sortArr[closest-1])
	}
	worst := 0
	worstStart := math.Abs(sortArr[0] - avg)
	worstEnd := math.Abs(sortArr[len(sortArr)-1] - avg)
	if worstStart > poorEnd {
		worst = 0
		fmt.Printf("最差评委是：评分%v的评委\n", sortArr[worst])
	} else if worstStart < worstEnd {
		worst = len(sortArr) - 1
		fmt.Printf("最差评委是：评分%v的评委\n", sortArr[worst])
	} else {
		fmt.Printf("最差评委是：评分%v、%v的评委\n", sortArr[0], sortArr[len(sortArr)-1])
	}

}
