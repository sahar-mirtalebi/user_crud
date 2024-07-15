package main

import (
	"fmt"
	"math/rand"
)

type UserService struct {
	repo *UserRepository
}

func newUserServise(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) createUser() {
	firstName, lastName, age, email := getUserInputs()

	user := UserData{
		userId:    rand.Intn(1000000),
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		email:     email,
	}
	service.repo.addUser(user)
	fmt.Printf("user %v created successfully\n", user)
}

func (service *UserService) updateUser() {
	var id int
	fmt.Println("please give me the id you want to update :")
	fmt.Scan(&id)

	for _, user := range service.repo.users {
		if user.userId == id {
			firstName, lastName, age, email := getUserInputs()

			updatedUser := UserData{
				userId:    rand.Intn(1000000),
				firstName: firstName,
				lastName:  lastName,
				age:       age,
				email:     email,
			}

			service.repo.updateUser(updatedUser)
			return
		}

	}
	fmt.Println("no such user")
}

func (service *UserService) deleteUser() {
	var id int
	fmt.Println("please give me the id you want to delete :")
	fmt.Scan(&id)

	service.repo.deleteUser(id)
}

func (service *UserService) retrieveAllUsers() {
	users := service.repo.retrieveAllUsers()
	for _, user := range users {
		fmt.Println(user)
	}
}

func (service *UserService) retrieveUserById() {
	var id int
	fmt.Println("please give me the id you want to retrieve : ")
	fmt.Scan(&id)

	user, found := service.repo.retrieveUserById(id)
	if found {
		fmt.Println(user)
	} else {
		fmt.Println("user not found")
	}
}
