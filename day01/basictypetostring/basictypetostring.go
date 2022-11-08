package main

import "fmt"

func main() {
	var num1 int = 99
	fmt.Println(num1)
	var str string
	var f float32 = 0.11
	var b bool = true
	var myChar byte = 'h'

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("b type %T b=%q\n", str, str)

	str = fmt.Sprintf("%f", f)
	fmt.Printf("b type %T b=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("b type %T b=%q\n", str, str)
}
