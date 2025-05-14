package process
import ( 
	"fmt"
	"net"
	"message"
	"encoding/json"
	"utils"
)

type UesrProcess struct {
	Conn net.Conn
}

//编写一个函数专门处理登录请求
func (this *UesrProcess) ServerProcessLogin(mes *message.Message) (err error){

	//1.从mes取出data反序列化
	var loginMes message.LoginMes 
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("serverProcessLogin json unmaishal fail, err=", err)
		return
	}
	
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200

	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请先注册"
	}

	//loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("serverProcessLogin loginResMes json maishal fail, err=", err)
		return
	}

	resMes.Data = string(data)	
	data, err = json.Marshal(resMes) 
	if err != nil {
		fmt.Println("serverProcessLogin resMes json maishal fail, err=", err)
		return
	}

	//创建Tranfer实例
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
    //发送
	err = tf.WritePkg(data)
	return
}
