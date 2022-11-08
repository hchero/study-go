package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @description:
 * @param {*[]int} arr
 * @param {int} insertNum
 * @return {*}
 */
func dichotomyInsert(arr *[]int, insertNum int) {
	/**
	* @description
	1、先增加元素
	2、再查找位置
	3、再挪数据
	* @return {*}
	*/
	*arr = append(*arr, insertNum)
	for i := 0; i <= len(*arr); i++ {

	}
	// 二分法 dichotomy binary find
	insertIndex := dichotomy(arr, insertNum)
	for i := len(*arr) - 1; i > insertIndex; i-- {
		(*arr)[i] = (*arr)[i-1]
	}
	(*arr)[insertIndex] = insertNum
}

func dichotomy(arr *[]int, insertNum int) int {
	var start int = 0
	var end int = len(*arr) - 1
	if insertNum < (*arr)[0] {
		return 0
	}
	if insertNum > (*arr)[len(*arr)-1] {
		return len(*arr)
	}
	for start < end {
		middle := (start + end) / 2
		if (*arr)[middle] == insertNum {
			return middle
		} else if (*arr)[middle] > insertNum {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return end
}

func main() {
	arr := []int{1, 2, 36, 49, 61, 100}
	fmt.Println(arr)
	rand.Seed(time.Now().Unix())
	insertNum := rand.Intn(100)
	fmt.Println("需要插入的数是；", insertNum)
	end := dichotomy(&arr, 56)
	fmt.Println("插入位置是：", end)
	dichotomyInsert(&arr, 56)
	fmt.Println("新数组是；", arr)
}
