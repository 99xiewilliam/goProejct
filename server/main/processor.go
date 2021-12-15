package main

import (
	"client/main/chatroom/common/message"
	process2 "client/main/chatroom/server/process"
	"client/main/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn  net.Conn
}
func (this *Processor) serverProcessMes (mes *message.Message) (err error) {

	//看看那是否能够接收到客户端发送群发的消息
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		//
		up := &process2.UserProcessor{
		 Conn : this.Conn,
		}
		err = up.ServerProcessLogin(this.Conn, mes)
	case message.RegisterMesType:
		//
		up := &process2.UserProcessor{
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)
		fmt.Println("++++++++++++++++")
		fmt.Println(err)
	case message.SmsMesType:
		//创建一个smsProcess实例完成转发群聊消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在")
	}
	return
}
//func process(conn net.Conn) {
//	defer conn.Close()
//
//	for {
//		mes, err := readPkg(conn)
//
//		if err != nil {
//			if err == io.EOF {
//				fmt.Println("客户端退出，服务器退出")
//				return
//			}else {
//				fmt.Println("readPkg err=", err)
//			}
//
//		}
//
//		//fmt.Println("mes=", mes)
//		err = serverProcessMes(conn, &mes)
//
//		if err != nil {
//			return
//		}
//
//	}
//}

func (this *Processor) process2() (err error) {
	for {
		tf := &utils.Transfer{Conn : this.Conn}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器退出")
				return err
			}else {
				fmt.Println("readPkg err=", err)
				return err
			}

		}

		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)

		if err != nil {
			return err
		}

	}
}
