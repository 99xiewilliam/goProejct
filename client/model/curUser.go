package model

import (
	"client/main/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
