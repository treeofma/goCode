package main
import(
	"fmt"
)


func main() {
	intChan := make(chan int, 3)
	intChan<- 100
	intChan<- 200

	close(intChan)//关闭不能在写入数据，但可以读数据


	//遍历
	intChan2 := make(chan int, 50)
	for i := 0; i < 50; i++ {
		intChan2<- i * 2
	}

	//遍历时关闭管道会正常便利数据否则会出现deadlock错误
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}


}