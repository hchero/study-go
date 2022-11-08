package main

import "fmt"

func main() {
	str := "hello word,你好，世界"
	slice := str[:]
	fmt.Println(slice)
	fmt.Printf("slice type = %T", slice)
	// arr := []byte(str)
	arr := []rune(str)
	arr[0] = '呀'
	fmt.Printf("%c", arr)
	fmt.Printf("slice type = %T", arr)

}
