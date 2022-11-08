package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = " hello你好 "
	fmt.Println("str len=", len(str))
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c\n", str[i])
	// }
	// r := []rune(str)
	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("字符=%c\n", r[i])
	// }

	str2 := strings.Trim(str, " ")
	fmt.Println("str2 len=", len(str2))
	fmt.Println("str = ", str, "\nstr2=", str2)
	str3 := strings.ToUpper(str)
	str4 := strings.ToLower(str3)
	fmt.Println("str3 = ", str3)
	fmt.Println("str4 = ", str4)
}
