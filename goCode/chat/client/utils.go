package main
import (
	"fmt"
	"net"
	"message"
	"encoding/binary"
	"encoding/json"
	"errors"
)
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)

	fmt.Println("读取服务端发送的数据")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("read err, err=", err) 
		//err = errors.New("read len err")
		return
	}
	fmt.Println("读到的buf=", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	n, err := conn.Read(buf[:pkgLen]) //从conn中读pkglen长放入buf中
	if n != int(pkgLen) || err != nil {
		fmt.Println("read data err, err=", err)
		err = errors.New("read data err")
		return
	}

	//反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	
	return 
}

func writePkg(conn net.Conn, data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4]) //从conn中读pkglen长放入buf中
	if n != 4 || err != nil {
		fmt.Println("writePkg write bytes err, err=", err)
		return
	}

	n, err = conn.Write(data)
	if n != int(pkgLen)  || err != nil {
		fmt.Println("writePkg write data err, err=", err)
		return
	}
	
	return 
}
