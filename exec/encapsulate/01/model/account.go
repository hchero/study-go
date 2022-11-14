package model

import "fmt"

type account struct {
	name     string
	balance  float64
	password string
}

func NewAccount() *account {
	return &account{}
}

func (a *account) SetAccount(name string) {
	if len(name) >= 6 && len(name) <= 10 {
		a.name = name
	} else {
		fmt.Println("账号长度在6-10之间")
	}
}

func (a *account) SetBalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("余额不能低于20")
	}
}

func (a *account) SetPassword(password string) {
	if len(password) == 6 {
		a.password = password
	} else {
		fmt.Println("密码必须是6位")
	}
}
