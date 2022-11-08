package main

import "fmt"

func main() {
	// 字符串遍历的两种方式
	// var str string = "hello"
	// fmt.Printf("第一种方式\n")
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n", str[i])
	// }

	// fmt.Printf("第二种方式\n")
	// for index, val := range str {
	// 	fmt.Printf("index=%d,val=%c \n", index, val)
	// }

	var str1 string = "hello,中国"
	str2 := []rune(str1) // 把str1转换为[]rune(切片)
	fmt.Printf("第一种方式\n")
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

}
