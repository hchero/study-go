/*
 * @Author: super zhou4xin3xin3520@163.com
 * @Date: 2022-11-17 16:30:22
 * @LastEditors: super
 * @Description:家庭收支明细软件
 */
package main

import "fmt"

var money int
var description string

func menu() {
	fmt.Println("----------家庭收支记账软件----------")
	fmt.Println("          1 收支明细")
	fmt.Println("          2 登记收入")
	fmt.Println("          3 登记支出")
	fmt.Println("          4 退出软件")
	fmt.Println("          请选择(1-4):")
}

func caseFunc(sf int) {
	switch sf {
	case 1:
		fmt.Println("收支\t账号余额\t收支金额\t说明")
		fmt.Println(description)

	case 2:
		var getMoney int
		fmt.Printf("本次收入金额：")
		fmt.Scanln(&getMoney)
		fmt.Println()
		var desc string
		fmt.Printf("本次收入说明：")
		fmt.Scanln(&desc)
		fmt.Println()
		money += getMoney
		currentDesc := "收入\t" + fmt.Sprintf("%d", money) + "\t\t" + fmt.Sprintf("%d", getMoney) + "\t\t" + desc + "\n"
		description += currentDesc

	case 3:
		var loseMoney int
		fmt.Printf("本次支出金额：")
		fmt.Scanln(&loseMoney)
		fmt.Println()
		var desc string
		fmt.Printf("本次支出说明：")
		fmt.Scanln(&desc)
		fmt.Println()
		money -= loseMoney
		currentDesc := "支出\t" + fmt.Sprintf("%d", money) + "\t\t" + fmt.Sprintf("%d", loseMoney) + "\t\t" + desc + "\n"
		description += currentDesc
	case 4:
	default:
		fmt.Println("功能不存在")
	}
}

func main() {

	money = 10000
	for {
		menu()
		var sf int
		fmt.Scanf("%d\n", &sf)
		if sf == 4 {
			break
		}
		caseFunc(sf)
	}
}
