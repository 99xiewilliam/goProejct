package main

import (
	"client/main/chatroom/client/process"
	"fmt"
)
var userId int
var userPwd string
var userName string

func main() {
	var key int
	//var loop = true
	//fmt.Println(loop)


	for true{
		fmt.Println("--------------welcome to login page------------")
		fmt.Println("\t\t\t 1 login chatroom")
		fmt.Println("\t\t\t 2 register user")
		fmt.Println("\t\t\t 3 exit system")
		fmt.Println("\t\t\t please select 1-3")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("login chatroom")
			fmt.Println("Please input id number")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Please input password")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("register user")
			fmt.Println("Please input user id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Please input user password:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("Please input user name:")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("exit system")
			//loop = false
		default:
			fmt.Println("input wrong")

		}
	}

	//if key == 1 {
	//
	//
	//	login(userId, userPwd)
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//}else {
	//	//	//fmt.Println("login success")
	//	//}
	//}else if key == 2 {
	//	fmt.Println("register logic.....")
	//}
}
