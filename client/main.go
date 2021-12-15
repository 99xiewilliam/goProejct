package main

import (
	"fmt"
)
var userId int
var userPwd string

func main() {
	var key int
	var loop = true
	fmt.Println(loop)

	for loop{
		fmt.Println("--------------welcome to login page------------")
		fmt.Println("\t\t\t 1 login chatroom")
		fmt.Println("\t\t\t 2 register user")
		fmt.Println("\t\t\t 3 exit system")
		fmt.Println("\t\t\t please select 1-3")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("login chatroom")
			loop = false
		case 2:
			fmt.Println("register user")
			loop = false
		case 3:
			fmt.Println("exit system")
			loop = false
		default:
			fmt.Println("input wrong")

		}
	}

	if key == 1 {
		fmt.Println("Please input id number")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("Please input password")
		fmt.Scanf("%s\n", &userPwd)

		login(userId, userPwd)
		//if err != nil {
		//	fmt.Println(err)
		//}else {
		//	//fmt.Println("login success")
		//}
	}else if key == 2 {
		fmt.Println("register logic.....")
	}
}
