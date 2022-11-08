package main

import "fmt"

func main() {
	fmt.Printf("请输入成绩\n")
	var result float32
	fmt.Scanf("%f", &result)
	fmt.Println("\n", result)
	fmt.Printf("\n 成绩 = %v\n", result)
	switch {
	case result >= 60 && result <= 100:
		fmt.Printf("合格")
	case result < 60 && result >= 0:
		fmt.Printf("不合格")
	case result > 100:
		fmt.Printf("成绩不能高于100")
	default:
		fmt.Printf("成绩不合法")
	}
}
