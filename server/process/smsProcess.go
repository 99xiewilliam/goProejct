package process

import (
	"client/main/chatroom/common/message"
	"client/main/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//.
}
//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端的onlineUsers map[int]*Userprocess
	//将消息转发取出

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)

	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里还需要过滤自己，即不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineuser(data, up.Conn)
	}

}

func (this *SmsProcess) SendMesToEachOnlineuser(data []byte, conn net.Conn) {

	//创建一个Transfer 实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
		return
	}
}