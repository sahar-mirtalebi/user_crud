package main

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
