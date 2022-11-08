package main

import "fmt"

func main() {
	fmt.Printf("请输入用户名\n")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("请输入密码\n")
	var password string
	fmt.Scanln(&password)
	if name == "张三" && password == "1234" {
		fmt.Printf("登录成功")
	} else {
		fmt.Printf("登录失败")
	}
}
