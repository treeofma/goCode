package main
import (
	"fmt"
	"net"
	"io"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		//循环接受客户端发送得数据
		//创建一个新的切片
		buf := make([]byte, 1024)

		//read阻塞等待客户端发来得数据
		n, err := conn.Read(buf)  	
		if err == io.EOF {
			fmt.Println("客户端退出 err=", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听......")

	//1.tcp表示使用协议， 127.0.0.1:9000表示本地监听90端口
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("监听失败，err=", err)
		return
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端链接....")
		
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		} else {
			fmt.Printf("Accept succ conn=%v, 客户端ip%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)

	}

}