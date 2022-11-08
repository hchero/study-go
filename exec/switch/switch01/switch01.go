package main

import "fmt"

func main() {
	var char byte

	fmt.Printf("请输入字符\n")
	fmt.Scanf("%c", &char)
	fmt.Printf("c type %T c=%v\n", char, char)
	switch char {
	case 'a', 'b', 'c', 'd', 'e':
		// 字符转换大写
		var change byte = char - 32
		fmt.Printf("%c", change)
	default:
		fmt.Printf("other")
	}
}
