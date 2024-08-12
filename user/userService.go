package user

type UserService struct {
	repo *UserRepository
}

func NewUserServise(repo *UserRepository) UserOperations {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(userData UserData) (uint, error) {

	user := &UserData{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Age:       userData.Age,
		Email:     userData.Email,
	}
	err := service.repo.AddUser(user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil

}

func (service *UserService) UpdateUser(updatedUserData UserData) error {
	return service.repo.UpdateUser(updatedUserData)
}

func (service *UserService) DeleteUser(userId uint) error {

	return service.repo.DeleteUser(userId)
}

func (service *UserService) RetrieveAllUsers(name, email string, page, limit int) ([]UserData, error) {
	var offset = (page - 1) * 2
	return service.repo.RetrieveAllUsers(name, email, offset, limit)

}

func (service *UserService) RetrieveUserById(userId uint) (UserData, error) {
	return service.repo.RetrieveUserById(userId)

}
