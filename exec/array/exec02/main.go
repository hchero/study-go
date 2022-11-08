package main

import "fmt"

/**
 * @description: 求数组最大值,并获取对应的下标,请求出一个数组的和和平均值,for-range
 * @return {*}
 */
func main() {
	var arr = [...]int{1, 39, 34, 21, 35, 3, 9}
	var max int
	var maxIndex int
	var total int
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 		maxIndex = i
	// 	}
	// 	total += arr[i]
	// }
	for i, v := range arr {
		if v > max {
			max = v
			maxIndex = i
		}
		total += v
	}

	vag := fmt.Sprintf("%.2f", float64(total)/float64(len(arr)))
	fmt.Printf("最大值%v的下标%v,平均值%v", max, maxIndex, vag)
}
