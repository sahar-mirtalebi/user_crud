package main

import "fmt"

func updateUser() {
	var updatedUserId int
	fmt.Println("please give me the user id you want to update :")
	fmt.Scan(&updatedUserId)
	success := false
	for i, user := range users {
		if user.userId == updatedUserId {
			fmt.Println("please give me your new information :")
			updatedData := getUserInputs()
			users[i].firstName = updatedData.firstName
			users[i].lastName = updatedData.lastName
			users[i].age = updatedData.age
			users[i].email = updatedData.email
			success = true
		}
	}
	if success {
		fmt.Println("user updated successfully")
	} else {
		fmt.Println("user not found")
	}
}
