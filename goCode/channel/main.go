package main

import (
	"fmt"
)

func main() {
	//1.创建存放3个int的管道
	var intChan chan int 
	intChan = make(chan int, 3)

	//2.写入数据
	intChan<- 10
	intChan<- 12

	//3.大小和容量
	fmt.Printf("intChan len is %d, cap is %d\n", len(intChan), cap(intChan))

	//4.从管道读取数据
	var num int
	num = <-intChan
	fmt.Println(num)
}