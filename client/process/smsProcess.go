package process

import (
	"client/main/chatroom/common/message"
	"client/main/chatroom/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//创建一个message
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建一个SmsMes实例

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal fail=", err.Error())
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}
