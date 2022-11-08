package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "1好"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Println("转换成功", num)
	}
}
