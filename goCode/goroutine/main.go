package main
import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("helloworld")
		time.Sleep(time.Second)   //睡眠1秒
	}
}
func main() {

	go test()   //开启一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)   //睡眠1秒
	}
}