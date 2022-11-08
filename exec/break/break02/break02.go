package main

import "fmt"

func main() {
	i := 1
	for {

		fmt.Printf("请输入用户名\n")
		var name string
		fmt.Scanln(&name)
		fmt.Printf("请输入密码\n")
		var password string
		fmt.Scanln(&password)
		if name == "张无忌" && password == "888" {
			fmt.Printf("登录成功\n")
			break
		} else {
			if i >= 3 {
				fmt.Printf("没有机会了\n")
				break
			} else {
				fmt.Printf("还有%v次机会\n", 3-i)
			}
		}
		i++
	}

}
