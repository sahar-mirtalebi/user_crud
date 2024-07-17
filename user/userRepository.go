package user

import "fmt"

type UserData struct {
	UserId    int
	FirstName string
	LastName  string
	Age       int
	Email     string
}

type UserRepository struct {
	users []UserData
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make([]UserData, 0)}
}

func (repo *UserRepository) addUser(userData UserData) {
	for _, user := range repo.users {
		if user.Email == userData.Email {
			return
		}
	}
	repo.users = append(repo.users, userData)
}

func (repo *UserRepository) deleteUser(id int) {
	for i, user := range repo.users {
		if user.UserId == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return
		}
	}
}

func (repo *UserRepository) updateUser(updateUser UserData) {
	for i, user := range repo.users {
		if user.UserId == updateUser.UserId {
			repo.users[i] = updateUser
			return
		}
	}
}

func (repo *UserRepository) retrieveAllUsers() []UserData {
	return repo.users
}

func (repo *UserRepository) retrieveUserById(id int) UserData {
	for _, user := range repo.users {
		if user.UserId == id {
			fmt.Println(user)
			return user
		}
	}
	fmt.Println("user not found")
	return UserData{}
}
