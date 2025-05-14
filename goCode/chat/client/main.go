package main
import (
	"fmt"
)

func main() {

	//接受用户选则
	var key int

	//判断是否显示菜单
	var loop bool =  true

	var userId int
	var userPwd string

	for loop {
		fmt.Println("-------------------------欢迎登录多人聊天室----------------------------------")
		fmt.Println("\t\t\t 1 登录")
		fmt.Println("\t\t\t 2 注册")
		fmt.Println("\t\t\t 3 退出")
		fmt.Println("\t\t\t 请选择(1-3)")


		fmt.Scanf("%d\n", &key)
		switch key {
			case 1:
				fmt.Println("登录聊天室")
				loop = false
			case 2:
				fmt.Println("注册用户")
				loop = false
			case 3:
				fmt.Println("退出登录")
				loop = false
			default:
				fmt.Println("输入有误，请选择(1-3)")
		}
	}


	//根据用户输入，显示新的提示信息
	if key == 1 {
		fmt.Println("请输入用户id:")
		fmt.Scanf("%d\n", &userId)

		fmt.Println("请输入用户密码:")
		fmt.Scanf("%s\n", &userPwd)

		err := login(userId, userPwd) 
		if err != nil {
			fmt.Println("登陆失败")
		} 

	} 
}