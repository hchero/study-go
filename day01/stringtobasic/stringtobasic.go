package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "1234"
	var num int64
	/*
		strconv.ParseInt(str) 函数会返回两个值(value int,err error)
		只想获取value int,不想获取 err,需要使用_忽略
	*/
	num, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("num=%d \n", num)
	var str1 string = "11.001"
	var num2 float64
	num2, _ = strconv.ParseFloat(str1, 64)
	fmt.Printf("num2 type %T num2=%f", num2, num2)

	// 字符串转不了整数,默认返回0
	var num3 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(num3, 10, 64)
	fmt.Printf("n3 type %T n3=%d", n3, n3)
}
