/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-14 07:56:28
 * @LastEditors: super
 * @Description:
 */
package main

import (
	"fmt"
	"go_code/project01/exec/encapsulate/01/model"
)

func main() {
	account := model.NewAccount()
	account.SetAccount("abcde111fsfafa")
	account.SetBalance(2.0)
	account.SetPassword("1456")
	fmt.Println(account)
}
