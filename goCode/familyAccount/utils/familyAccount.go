package utils 
import (
	"fmt"
)

type FamilyAccount struct {
	//用来接收用户输入
	key string 

	//声明一个变量控制是否退出for循环
	loop bool

	//定义账户余额
	balance float64
	
	//每次收支金额
	money float64
	//每次收支说明
	note string
	//收支详情
	details string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance: 10000.0,
		money : 0.0,
		note : "",
		details: "收支\t账户金额\t收支金额\t说明\n",
	}
}

func (this *FamilyAccount) ShowMenu() {
	for {
		fmt.Println("----------家庭收支记账软件----------\n")
		fmt.Println("           1. 收支明细")
		fmt.Println("           2. 登记收入")
		fmt.Println("           3. 登记支出")
		fmt.Println("           4. 退出软件")

		fmt.Println("请输入选择：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("输入错误，请输入1-4")
		}
		if this.loop == false {
			break
		}

	}

	
}	

func (this *FamilyAccount) showDetails() {
	fmt.Println("----------当前收支明细记录----------")
	if this.details == "收支\t账户金额\t收支金额\t说明\n" {
		fmt.Println("当前没有收支明细记录...\n")
	} else {
		fmt.Println(this.details)
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("----------登记收入----------")
	fmt.Println("请输入收入金额：")
	fmt.Scanln(&this.money)
	fmt.Println("请输入收入说明：")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", this.balance, this.money, this.note)
}

func (this *FamilyAccount) pay() {
	fmt.Println("----------登记支出----------")
	fmt.Println("请输入支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	fmt.Println("请输入支出说明：")
	fmt.Scanln(&this.note)
	this.balance -= this.money
	this.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", this.balance, this.money, this.note)
}

func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}
