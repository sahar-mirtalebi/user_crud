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
	users     map[int]UserData
	userEmail map[string]int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:     make(map[int]UserData),
		userEmail: make(map[string]int),
	}
}

func (repo *UserRepository) addUser(userData UserData) {
	_, exists := repo.userEmail[userData.Email]
	if !exists {
		repo.users[userData.UserId] = userData
		repo.userEmail[userData.Email] = userData.UserId
	}
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

	user, exists := repo.users[userId]

	if exists {
		fmt.Println(user)
		return user
	} else {
		fmt.Println("User does not exist.")
		return UserData{}
	}

}
