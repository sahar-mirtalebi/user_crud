package main

import (
	"fmt"
	"math/rand"
	"time"
	"user_CRUD/user"
)

func main() {

	repo := user.NewUserRepository()
	service := user.NewUserServise(repo)

	var lestItemId int

	for i := 0; i < 100000; i++ {
		userData := randomData()

		lestItemId = service.CreateUser(userData)

	}

	startTime := time.Now()

	service.RetrieveUserById(lestItemId)

	endTime := time.Now()

	duration := endTime.Sub(startTime)

	fmt.Println(duration)

}

func generateRandomString(length int) string {
	const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)

	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}

func randomData() user.UserData {

	var firstNames = []string{"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda", "William", "Elizabeth"}
	var lastNames = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	age := rand.Intn(100)
	email := generateRandomString(10) + "@example.com"

	return user.UserData{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Email:     email,
	}
}
