package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	// 打印会增加程序耗时
	// fmt.Println(str)
}
func main() {
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test()耗时%v秒", end-start)
}
