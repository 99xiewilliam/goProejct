package model

import (
	"client/main/chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"

)

var (
	MyUserDao *UserDao
)



//使用工厂模式，创建一个UserDap实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error){
	fmt.Println("where is id")
	fmt.Println(id)
	res, err := redis.String(conn.Do("HGet", "users", id))
	fmt.Println(err == nil)

	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return

}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error){
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)

	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return

}

func (this *UserDao) Register(user *message.User) (err error){
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	fmt.Println("Resgister")
	fmt.Println(err)

	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	fmt.Println("------------------")
	fmt.Println(user)

	data, err := json.Marshal(user)

	if err != nil {
		return
	}
	fmt.Println(string(data))
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}

	return


}

