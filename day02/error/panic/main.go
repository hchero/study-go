package main

import (
	"fmt"
)

func readConf(name string) (err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	var a int = 1
	var b int = 0
	i := a / b
	fmt.Println(i)
	return nil
	// if name == "config.ini" {
	// 	return nil
	// } else {
	// 	return errors.New("读取文件错误... ")
	// }
}
func test2() {
	err := readConf("config.ini")
	if err != nil {
		// 如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test2()继续执行")
}
func main() {
	test2()
	fmt.Println("继续执行")
}
