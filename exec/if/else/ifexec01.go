package main

import "fmt"

func main() {
	fmt.Printf("请输入身高，财富，是否帅，用空格隔开：\n")
	var length int
	var money int
	var isHandsome bool
	fmt.Scanf("%d %d %t", &length, &money, &isHandsome)
	if length > 180 && money > 10000000 && isHandsome {
		fmt.Printf("我一定要嫁给他!!!")
	} else if length > 180 || money > 10000000 || isHandsome {
		fmt.Printf("嫁吧，比上不足，比下有余")
	} else if !(length > 180 || money > 10000000 || isHandsome) {
		fmt.Printf("不嫁")
	}
}
