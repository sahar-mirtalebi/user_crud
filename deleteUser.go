package main

import "fmt"

func deleteUser() {
	var deleteId int
	fmt.Println("please give me the user id you want to delete :")
	fmt.Scan(&deleteId)

	for i, user := range users {
		if user.userId == deleteId {
			users = append(users[:i], users[i+1:]...)
		}
	}
}
