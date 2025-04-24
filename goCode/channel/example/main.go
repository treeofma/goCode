package main
import(
	"fmt"
)


func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan<- i
	}
	close(intChan)
}

func readDate(intChan chan int, exitChan chan bool) {
	for  {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	exitChan<- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readDate(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		} 
	}
	
}