package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	ResiterMesType = "ResiterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}


//登录消息
type LoginMes struct {
	UserId int
	UserPwd string
	UserName string
}

type LoginResMes struct {
	Code int  //返回的状态码，500：未注册  200：登陆成功
	Error string  //返回错误信息
}

type ResiterMes struct {

}