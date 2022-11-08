package main

import "fmt"

func binaryFind(arr *[]int, left int, right int, findVal int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2
	fmt.Println("middle=", middle, (*arr)[middle])
	if (*arr)[middle] > findVal {
		fmt.Println(left, middle-1)
		binaryFind(arr, left, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		fmt.Println(middle+1, right)
		binaryFind(arr, middle+1, right, findVal)
	} else {
		fmt.Println("找到了")
		return middle
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	findIndex := binaryFind(&arr, 0, len(arr)-1, 4)
	fmt.Printf("下标是：%v", findIndex)

}
