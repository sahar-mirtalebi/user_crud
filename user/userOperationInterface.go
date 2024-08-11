package user

type UserOperations interface {
	CreateUser(userData UserData) (uint, error)
	DeleteUser(userId uint) error
	UpdateUser(userData UserData) error
	RetrieveAllUsers(name, email string, page, limit int) ([]UserData, error)
	RetrieveUserById(userId uint) (UserData, error)
}
