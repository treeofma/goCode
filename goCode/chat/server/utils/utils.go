package utils
import (
	"fmt"
	"net"
	"message"
	"encoding/binary"
	"encoding/json"
	"errors"
)


//把方法关联到结构体
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("read err, err=", err) 
		//err = errors.New("read len err")
		return
	}
	fmt.Println("读到的buf=", this.Buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen]) //从conn中读pkglen长放入buf中
	if n != int(pkgLen) || err != nil {
		fmt.Println("read data err, err=", err)
		err = errors.New("read data err")
		return
	}

	//反序列化
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	
	return 
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4]) //从conn中读pkglen长放入buf中
	if n != 4 || err != nil {
		fmt.Println("writePkg write bytes err, err=", err)
		return
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen)  || err != nil {
		fmt.Println("writePkg write data err, err=", err)
		return
	}
	
	return 
}