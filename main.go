package main

import (
	"fmt"
	"math/rand"
)

var users = make([]UserData, 0)

type UserData struct {
	userId    int
	firstName string
	lastName  string
	age       int
	email     string
}

func main() {

	createUser()

	createUser()

	createUser()

	retrieveAllUsers()

	retrieveUserById()

	updateUser()

	deleteUser()

	retrieveAllUsers()

}

func getUserInputs() UserData {
	var firstName string
	var lastName string
	var age int
	var email string

	fmt.Println("enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("enter your age : ")
	fmt.Scan(&age)

	fmt.Println("enter your email : ")
	fmt.Scan(&email)

	return UserData{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		email:     email,
	}
}

func createUser() (UserData, []UserData) {

	data := getUserInputs()

	var userId = rand.Intn(1000000)

	data.userId = userId
	users = append(users, data)
	fmt.Printf("user %v created successfully\n", data)
	return data, users
}

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
