package user

import (
	"fmt"
	"math/rand"
)

type UserService struct {
	repo *UserRepository
}

func NewUserServise(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(userData UserData) {

	user := UserData{
		UserId:    rand.Intn(1000000),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Age:       userData.Age,
		Email:     userData.Email,
	}
	service.repo.addUser(user)
	fmt.Printf("user %v created successfully\n", user)
}

func (service *UserService) UpdateUser(userId int, updatedUserData UserData) {

	user := UserData{
		UserId:    userId,
		FirstName: updatedUserData.FirstName,
		LastName:  updatedUserData.LastName,
		Age:       updatedUserData.Age,
		Email:     updatedUserData.Email,
	}

	service.repo.updateUser(user)

}

func (service *UserService) DeleteUser(userId int) {

	service.repo.deleteUser(userId)
}

func (service *UserService) RetrieveAllUsers() {
	users := service.repo.retrieveAllUsers()
	for _, user := range users {
		fmt.Println(user)
	}
}

func (service *UserService) RetrieveUserById(userId int) {
	user, found := service.repo.retrieveUserById(userId)
	if found {
		fmt.Println(user)
	} else {
		fmt.Println("user not found")
	}
}
