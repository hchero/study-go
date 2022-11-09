/*
* @Author: super zhou4xin3xin3520@163.com
* @Date: 2022-11-09 16:00:04
  - @LastEditors: super
  - @LastEditTime: 2022-11-09 17:43:39

* @Description:
1、使用map[string]map[string]string的map类型
2、key:表示用户名，是唯一的，不吭重复
3、如果某个用户名存在，就将其密码修改“888888”，如果不存在就增加这个用户信息(包括昵称nickname和密码pwd)
4、编写一个函数modifyUser(users map[string]map[string]string,name string)完成上述功能
*/
package main

import "fmt"

/**
 * @description: 修改指定用户的密码，不存在则增加这个用户的信息
 * @param {map[string]map[string]string} users
 * @return {*}
 */
func modifyUser(users map[string]map[string]string, name string) map[string]map[string]string {
	_, findRes := users[name]
	if findRes {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = map[string]string{
			"nickname": name,
			"pwd":      "888888",
		}
	}
	return users
}

func main() {
	var user map[string]map[string]string
	user = make(map[string]map[string]string)
	user["小明"] = map[string]string{"nickname": "小明", "pwd": "18"}
	user["小白"] = map[string]string{"nickname": "小白", "pwd": "18"}

	fmt.Println(user)
	user = modifyUser(user, "小白")
	user = modifyUser(user, "小花")
	fmt.Println(user)

}
