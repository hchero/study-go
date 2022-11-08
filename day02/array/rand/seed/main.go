package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机生成5个数,并将其反转打印
1. 生成随机数,rand.Intn()函数
2. 反转打印,交互的次数是 len / 2, 倒数第一个数和第一个元素交换,导数第二个元素和第二个元素交换
*/
func main() {
	var arr [9]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) // 0-100的随机数
	}
	fmt.Println(arr)
	fmt.Println("反转处理")
	for i := 0; i < len(arr)/2; i++ {
		changeNum := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = changeNum
	}
	fmt.Println(arr)
}
