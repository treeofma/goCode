package main 
import (
	"fmt"
	"net"
	"message"
	"process"
	"io"
	"utils"
)

type Processer struct {
	Conn net.Conn
}

//根据客户端发送的消息不同，决定调用哪个函数来处理
func (this * Processer) serverProcessMes(mes *message.Message) (err error){

	switch mes.Type {
		case message.LoginMesType:
			//登录
			userP := &process.UesrProcess {
				Conn : this.Conn,
			}
			err = userP.ServerProcessLogin(mes)
		case message.ResiterMesType:
			//注册：
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return

}


func (this * Processer) process2() (err error){
	for {

		tf := &utils.Transfer {
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭")
				return err
			} else {
				fmt.Println("in func Process readPkg err, err=", err)
				return err
			}
			
		}
	
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
		//fmt.Println("mes=", mes)
	}
}
