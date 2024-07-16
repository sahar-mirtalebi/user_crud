package main

import (
	"user_CRUD/user"
)

func main() {

	repo := user.NewUserRepository()
	service := user.NewUserServise(repo)

	userData := user.UserData{
		FirstName: "sahar",
		LastName:  "mirtalebi",
		Age:       32,
		Email:     "ss@gmail.com",
	}

	service.CreateUser(userData)
	service.CreateUser(userData)
	service.RetrieveAllUsers()

}
