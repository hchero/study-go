package main

import "fmt"

/*
创建一个byte类型的26个元素的数组,分别放置A-Z.使用for循环访问所有元素并打印出来,字符运算 'A' + 1 = 'B'
*/
func main() {
	var charArr [26]byte
	for i := 0; i < len(charArr); i++ {
		charArr[i] = 'A' + byte(i)
	}
	fmt.Printf("%c ", charArr)
}
