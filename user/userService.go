package user

import (
	"math/rand"
)

type UserService struct {
	repo *UserRepository
}

func NewUserServise(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(userData UserData) int {

	user := UserData{
		UserId:    rand.Intn(1000000),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Age:       userData.Age,
		Email:     userData.Email,
	}
	service.repo.addUser(user)
	return user.UserId
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

func (service *UserService) RetrieveAllUsers() map[int]UserData {
	return service.repo.retrieveAllUsers()

}

func (service *UserService) RetrieveUserById(userId int) UserData {
	return service.repo.retrieveUserById(userId)

}
