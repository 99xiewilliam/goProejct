package process

import (
	"client/main/chatroom/client/utils"
	"client/main/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess) Register(userId int,
	userPwd string, userName string) (err error)  {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes

	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)

	if err != nil {
		fmt.Println(err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)

	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}


	tf := &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送消息错误 err=", err)
	}
	fmt.Println(string(data))

	mes, err = tf.ReadPkg()
	fmt.Println(err)

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	var registerResMes message.RegisterResMes
	fmt.Println(mes.Data)

	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return

}

func (this *UserProcess)Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId=%d userPwd=%s\n", userId, userPwd)
	//return nil

	conn, err := net.Dial("tcp", "localhost:8889")

	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()


	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)

	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)

	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n ,err := conn.Write(buf[:4])

	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err=", err)
		return
	}

	//fmt.Printf("客户端，发送消息的长度=%d,内容=%s\n", len(data), string(data))
	_ ,err = conn.Write(data)

	if err != nil {
		fmt.Println("conn.Write(data) fail err=", err)
		return
	}

	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠20秒")

	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([] byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//fmt.Println("登录成功")
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}


	return

}



