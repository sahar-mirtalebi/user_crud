package user

import (
	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email" gorm:"unique"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) addUser(userData *UserData) error {
	return repo.db.Create(&userData).Error
}

func (repo *UserRepository) deleteUser(userId uint) error {
	return repo.db.Delete(&UserData{}, userId).Error
}

func (repo *UserRepository) updateUser(userData UserData) error {
	return repo.db.Save(&userData).Error
}

func (repo *UserRepository) retrieveAllUsers() ([]UserData, error) {
	var users []UserData
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *UserRepository) retrieveUserById(userId uint) (UserData, error) {
	var user UserData
	err := repo.db.First(&user, userId).Error
	return user, err
}
