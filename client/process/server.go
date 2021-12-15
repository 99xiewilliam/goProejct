package process

import (
	"client/main/chatroom/common/message"
	"client/main/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("-----------恭喜xxx登录成功------------")
	fmt.Println("-----------1.显示在线用户列表----------")
	fmt.Println("-----------2.发送消息-----------------")
	fmt.Println("-----------3.信息列表-----------------")
	fmt.Println("-----------4.退出系统-----------------")
	fmt.Println("请选择（1-4）：")
	var key int
	var content string
	//因为，我们总会用到Smsprocess实例，因此定义在外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)

	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说什么")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出系统...")
		os.Exit(0)
	default:
		fmt.Println("输入选项不正确")
	}
}

func serverProcessMes(Conn net.Conn) {
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		//fmt.Printf("mes=%v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)

			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType://有人群发消息
			outputGroupMes(&mes)

		default:
			fmt.Println("服务端返回未知消息类型。。。")
		}

	}
}
