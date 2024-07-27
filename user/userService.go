package user

import "gorm.io/gorm"

type UserService struct {
	repo *UserRepository
}

func NewUserServise(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(userData UserData) (uint, error) {

	user := &UserData{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Age:       userData.Age,
		Email:     userData.Email,
	}
	err := service.repo.addUser(user)

	if err != nil {
		return 0, err
	} else {
		return user.ID, nil
	}
}

func (service *UserService) UpdateUser(userId uint, updatedUserData UserData) error {

	user := UserData{
		Model:     gorm.Model{ID: userId},
		FirstName: updatedUserData.FirstName,
		LastName:  updatedUserData.LastName,
		Age:       updatedUserData.Age,
		Email:     updatedUserData.Email,
	}

	return service.repo.updateUser(user)
}

func (service *UserService) DeleteUser(userId uint) error {

	return service.repo.deleteUser(userId)
}

func (service *UserService) RetrieveAllUsers(page int, limit int) ([]UserData, error) {
	var offset = (page - 1) * 2
	return service.repo.retrieveAllUsers(offset, limit)

}

func (service *UserService) RetrieveUserById(userId uint) (UserData, error) {
	return service.repo.retrieveUserById(userId)

}
