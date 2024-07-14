package main

import "fmt"

func main() {

	repo := newUserRepository()
	service := newUserServise(repo)

	service.createUser()
	service.createUser()

	service.retrieveAllUsers()

	service.retrieveUserById()

	service.updateUser()

	service.deleteUser()

	service.retrieveAllUsers()

}

func getUserInputs() (string, string, int, string) {
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

	return firstName, lastName, age, email
}
