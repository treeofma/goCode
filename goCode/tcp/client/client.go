package main
import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)


func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("链接失败，err=\n", err)
		return
	}
	fmt.Println("成功 conn=", conn)
    
	for {
		//发送数据
		reader := bufio.NewReader(os.Stdin)

		//从终端读取i一行并发给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
			return
		}

		//处理line去掉\n
		ex := strings.Trim(line, " \r\n")
		if ex == "exit" {
			fmt.Println("客户端退出\n")
			break
		}
		n, err := conn.Write([]byte(line)) //n代表写了多少字节
		if err != nil {
			fmt.Println("write err=", err)
		}

		fmt.Printf("客户端发送了%d字节数据\n", n)
	}
	

}