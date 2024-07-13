package main

import "fmt"

func retrieveAllUsers() {
	fmt.Println(users)
}

func retrieveUserById() {
	var retrieveId int
	fmt.Println("please give me the id you want to retrieve :")
	fmt.Scan(&retrieveId)
	retrieve := false
	for _, user := range users {
		if user.userId == retrieveId {
			fmt.Println(user)
			retrieve = true
		}
	}
	if !retrieve {
		fmt.Println("no such user")
	}
}
