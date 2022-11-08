package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机生成10个整数（1-100的范围），保存到数组，并倒叙打印以及求平均值、最大值、最小值的下标，并查找里面是否有55
*/

/*
倒叙打印
*/
func sortPrint(arr [10]int) {
	// 倒叙打印
	var length int = len(arr)
	for i := 0; i < length; i++ {
		fmt.Println(arr[length-1-i])
	}
}

/*
返回平均值
*/
func avg(arr [10]int) float32 {
	var avg float32
	var total int
	for i := 0; i < 10; i++ {
		total += arr[i]
	}
	avg = float32(total / 10)
	return avg
}

/*
返回最大最小值
*/
func maxAndMin(arr [10]int) (int, int) {
	var max int = arr[0]
	var min int = arr[0]
	var maxIndex int = 0
	var minIndex int = 0
	for i := 1; i < 10; i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i

		} else if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	return maxIndex, minIndex
}

/*
查找某个数是否存在
*/
func findNum(arr [10]int, num int) bool {
	for i := 0; i <= 10/2; i++ {
		if arr[i] == num {
			return true
		} else if arr[10-1-i] == num {
			return true
		}
	}
	return false
}

func main() {
	var numArray [10]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(numArray); i++ {
		numArray[i] = rand.Intn(100)
	}
	fmt.Println(numArray)
	sortPrint(numArray)
	avg := avg(numArray)
	fmt.Println("平均值=", avg)
	max, min := maxAndMin(numArray)
	fmt.Println("最大值下标=", max, "\t最小值下标=", min)
	isHas := findNum(numArray, 55)
	if isHas {
		fmt.Println("55存在")
	} else {
		fmt.Println("55不存在")
	}
}
