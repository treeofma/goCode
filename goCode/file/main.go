package main 
import (
	"fmt"
	"os"
)

func main () {
	fmt.Println("面向对象")
	
	//打开文件
	file, err := os.Open("/home/ma/goCode/account/testMyAccount.go")

	if err != nil {
		fmt.Println("open file error=", err)
	} else {
		fmt.Println("file = %v", file)
	} 
	err = file.Close()



}