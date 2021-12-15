package process

import (
	"client/main/chatroom/common/message"
	"client/main/chatroom/server/model"
	"client/main/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	//字段？
	Conn net.Conn

	UserId int
}

func (this *UserProcessor) NotifyOthersOnLineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcessor) NotifyMeOnline(userid int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userid
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
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

	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)

	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}


}
func (this *UserProcessor) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte (mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	fmt.Println("************************")

	err = model.MyUserDao.Register(&registerMes.User)
	fmt.Println("=====+++++++========")
	fmt.Println(err)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	}else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

func (this *UserProcessor) ServerProcessLogin(conn net.Conn, mes * message.Message) (err error)  {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte (mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginMesType

	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505
			loginResMes.Error = "内部服务器错误..."
		}

	} else {
		loginResMes.Code = 200


		this.UserId = loginMes.UserId

		userMgr.AddOnLineUser(this)
		this.NotifyOthersOnLineUser(loginMes.UserId)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user,"login success")
	}



	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200
	//}else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在"
	//
	//}
	data, err := json.Marshal(loginResMes)

	if err != nil {
		fmt.Println("json Marshal fail err=", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

//func (this *UserProcessor) serverProcessMes (conn net.Conn, mes *message.Message) (err error) {
//
//	mes.Type = message.LoginMesType
//	switch mes.Type {
//	case message.LoginMesType:
//		//
//		err = this.serverProcessLogin(conn, mes)
//	case message.RegisterMesType:
//		//
//	default:
//		fmt.Println("消息类型不存在")
//	}
//	return
//}
