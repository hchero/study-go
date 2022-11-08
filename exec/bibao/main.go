package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 如果 name没有指定后缀，则加上，否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}
func main() {
	f2 := makeSuffix(".png")
	fmt.Println("文件名处理后=", f2("bird"))
}
