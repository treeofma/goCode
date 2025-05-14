package main
import (
	"fmt"
	"net"
)


func processListen(conn net.Conn) {

	defer conn.Close()


	processer := &Processer {
		Conn : conn,
	}

	err := processer.process2()
	if err != nil {
		fmt.Println("携程通信错位, err=", err)
		return
	}
}

func main() {
	//
	fmt.Println("服务器在8889端口监听")

	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听失败，err=", err)
		return
	}
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
		go processListen(conn)

	}
}
