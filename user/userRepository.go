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
	users map[int]UserData
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[int]UserData)}
}

func (repo *UserRepository) addUser(userData UserData) {
	for _, user := range repo.users {
		if user.Email == userData.Email {
			return
		}
	}
	repo.users[userData.UserId] = userData
}

func (repo *UserRepository) deleteUser(userId int) {
	for id := range repo.users {
		if id == userId {
			delete(repo.users, userId)
			return
		}
	}
}

func (repo *UserRepository) updateUser(updateUser UserData) {
	for id := range repo.users {
		if id == updateUser.UserId {
			repo.users[updateUser.UserId] = updateUser
			return
		}
	}
}

func (repo *UserRepository) retrieveAllUsers() map[int]UserData {
	return repo.users
}

func (repo *UserRepository) retrieveUserById(userId int) UserData {
	for id := range repo.users {
		if id == userId {
			fmt.Println(repo.users[userId])
			return repo.users[userId]
		}
	}
	fmt.Println("user not found")
	return UserData{}
}
