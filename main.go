package main

import (
	"log"
	"net/http"
	"user_CRUD/user"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=admin password=sahar223010 dbname=userdata port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user.UserData{})

	repo := user.NewUserRepository(db)
	service := user.NewUserServise(repo)
	handler := user.NewUserHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", handler.RetrieveAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handler.RetrieveUserById).Methods("GET")
	router.HandleFunc("/upload", handler.ImportUsersFromExcel).Methods("POST")

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

	// var lastItemId uint

	// for i := 0; i < 100000; i++ {
	// 	userData := randomData()
	// 	userId, err := service.CreateUser(userData)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	lastItemId = userId
	// }

	// startTime := time.Now()

	// service.RetrieveUserById(lastItemId)

	// endTime := time.Now()

	// duration := endTime.Sub(startTime)

	// fmt.Println(duration)

}

// func generateRandomString(length int) string {
// 	const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	result := make([]byte, length)

// 	for i := range result {
// 		result[i] = letter[rand.Intn(len(letter))]
// 	}
// 	return string(result)
// }

// func randomData() user.UserData {

// 	var firstNames = []string{"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael", "Linda", "William", "Elizabeth"}
// 	var lastNames = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}

// 	firstName := firstNames[rand.Intn(len(firstNames))]
// 	lastName := lastNames[rand.Intn(len(lastNames))]
// 	age := rand.Intn(100)
// 	email := generateRandomString(10) + "@example.com"

// 	return user.UserData{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Age:       age,
// 		Email:     email,
// 	}
// }
