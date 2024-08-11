package user

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
)

type UserHandler struct {
	operations UserOperations
}

func NewUserHandler(operations UserOperations) *UserHandler {
	return &UserHandler{operations: operations}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user UserData

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err := h.operations.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]uint{"id": userId})
}

func (h *UserHandler) ImportUsersFromExcel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		http.Error(w, "Failed to create temp file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Failed to copy file", http.StatusInternalServerError)
		return
	}
	tempFile.Close()

	f, err := excelize.OpenFile(tempFile.Name())
	if err != nil {
		http.Error(w, "Failed to open Excel file", http.StatusInternalServerError)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		http.Error(w, "Failed to get rows", http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	for _, row := range rows {
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
			age, _ := strconv.Atoi(row[2])
			user := UserData{
				FirstName: row[0],
				LastName:  row[1],
				Age:       age,
				Email:     row[3],
			}
			h.operations.CreateUser(user)
		}(row)

	}
	wg.Wait()
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user UserData
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = uint(userId)
	err = h.operations.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.operations.DeleteUser(uint(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) RetrieveAllUsers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	email := queryParams.Get("email")
	pageStr := queryParams.Get("page")
	limitStr := queryParams.Get("limit")

	page := 1
	limit := 10

	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	users, err := h.operations.RetrieveAllUsers(name, email, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) RetrieveUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user UserData
	user, err = h.operations.RetrieveUserById(uint(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
