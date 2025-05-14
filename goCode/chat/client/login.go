package main
import (
	"fmt"
	"encoding/json"
	"encoding/binary"
	"net"
	"message"
)

func login(userId int, userPwd string) (err error) {
	
	//连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接失败 err=", err) 
		return
	}

	//延时关闭
	defer conn.Close() 

	//通过conn发给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//序列化loginMes
	data, err := json.Marshal(loginMes)  //这里data是byte切片
	if err != nil {
		fmt.Println("loginMes json err, err=", err)
		return 
	}

	mes.Data = string(data)

	//序列化message
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json err, err=", err)
		return 
	}

	//发送长度，先获取data长度在转成一个byte切片，使用binary包数字转字节
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) 

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("len write err, err=", err)
		return
	}

	fmt.Printf("客户端发送消息长度成功, 长度为%d, 内容为%v\n", len(data), string(data))
	//time.Sleep(time.Second*5)


	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("data write err, err=", err)
		return
	}
	
	//处理服务器返回消息
	mes, err = readPkg(conn) 
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
	
}