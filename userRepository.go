package main

type UserData struct {
	userId    int
	firstName string
	lastName  string
	age       int
	email     string
}

type UserRepository struct {
	users []UserData
}

func newUserRepository() *UserRepository {
	return &UserRepository{users: make([]UserData, 0)}
}

func (repo *UserRepository) addUser(user UserData) {
	repo.users = append(repo.users, user)
}

func (repo *UserRepository) deleteUser(id int) {
	for i, user := range repo.users {
		if user.userId == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return
		}
	}
}

func (repo *UserRepository) updateUser(updateUser UserData) {
	for i, user := range repo.users {
		if user.userId == updateUser.userId {
			repo.users[i] = updateUser
			return
		}
	}
}

func (repo *UserRepository) retrieveAllUsers() []UserData {
	return repo.users
}

func (repo *UserRepository) retrieveUserById(id int) (UserData, bool) {
	for _, user := range repo.users {
		if user.userId == id {
			return user, true
		}
	}
	return UserData{}, false
}
