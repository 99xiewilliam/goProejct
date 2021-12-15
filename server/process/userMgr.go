package process

import (
	"fmt"
)
var(
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcessor
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

func (this *UserMgr) AddOnLineUser(up *UserProcessor) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DelOnLineUser(userId int) {
	delete(this.onlineUsers, userId)
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcessor {
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcessor, err error){
	up, ok := this.onlineUsers[userId]

	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}