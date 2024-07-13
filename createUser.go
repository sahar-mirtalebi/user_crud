package main

import (
	"fmt"
	"math/rand"
)

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
