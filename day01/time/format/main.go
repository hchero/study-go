package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 格式化第一种昂是
	var year = time.Now().Year()
	var month = time.Now().Month()
	var day = time.Now().Day()
	var hour = time.Now().Hour()
	var minute = time.Now().Minute()
	var second = time.Now().Second()
	fmt.Println(year, '-', month, '-', day, ' ', hour, ':', minute, ':', second)
	// 格式化第二种方式
	// var dateStr = now.Format("2006-01-02 15:04:05")
	var dateStr = now.Format("2006-01-02 15:04:05")
	fmt.Println(dateStr)
}
