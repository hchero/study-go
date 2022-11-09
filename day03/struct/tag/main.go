/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-09 17:50:52
 * @LastEditors: super
 * @LastEditTime: 2022-11-09 17:55:37
 * @Description: 标签 Tag
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	Monster := Monster{"牛魔王", 500, "芭蕉扇"}
	jsonStr, err := json.Marshal(Monster)
	if err != nil {
		fmt.Println("json处理失败", err)
	} else {
		fmt.Println("jsonStr", string(jsonStr))
	}

}
