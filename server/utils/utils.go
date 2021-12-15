package utils

import (
	"client/main/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//this.Buf = make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("err")
		return
	}
	//fmt.Println("读到的buf=", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)

	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return

}

func (this *Transfer) WritePkg(data []byte) (err error){
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])

	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	n, err = this.Conn.Write(data)

	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return

}
