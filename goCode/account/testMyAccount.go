package main

import (
	"fmt"
)

func main() {
	
	//用来接收用户输入
	var key string = ""

	//声明一个变量控制是否退出for循环
	var loop bool = true

	//定义账户余额
	var balance float64 = 10000.0
	
	//每次收支金额
	money := 0.0
	//每次收支说明
	note := ""
	//收支详情
	details := "收支\t账户金额\t收支金额\t说明\n"

	//显示主菜单
	for {
		fmt.Println("----------家庭收支记账软件----------\n")
		fmt.Println("           1. 收支明细")
		fmt.Println("           2. 登记收入")
		fmt.Println("           3. 登记支出")
		fmt.Println("           4. 退出软件")
		fmt.Println("           请选择(1-4):")
		
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----------当前收支明细记录----------")
			if details == "收支\t账户金额\t收支金额\t说明\n" {
				fmt.Println("当前没有收支明细记录...\n")
			} else {
				fmt.Println(details)
			}
		case "2":
			fmt.Println("----------登记收入----------")
			fmt.Println("请输入收入金额：")
			fmt.Scanln(&money)
			fmt.Println("请输入收入说明：")
			fmt.Scanln(&note)
			balance += money
			details += fmt.Sprintf("收入\t%v\t%v\t%v\n", balance, money, note)
		case "3":
			fmt.Println("----------登记支出----------")
			fmt.Println("请输入支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			fmt.Println("请输入支出说明：")
			fmt.Scanln(&note)
			balance -= money
			details += fmt.Sprintf("支出\t%v\t%v\t%v\n", balance, money, note)
		case "4":
			fmt.Println("----------你确定要退出吗？(y/n)----------")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入错误，请重新输入(y/n)")
			}
			if choice == "y" {
				loop = false
				break
			}
		default:
			fmt.Println("输入错误，请重新输入")
		}

		if loop == false {
			break
		}

	}
	fmt.Println("你已退出家庭记账软件...")
}