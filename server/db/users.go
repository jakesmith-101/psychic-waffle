package db

import (
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Nickname     string    `json:"nickname"`
	PasswordHash string    `json:"password"`
	RoleID       string    `json:"roleid"`
	AuthToken    string    `json:"authtoken"`
	CreatedAt    time.Time `json:"createdat"`
	UpdatedAt    time.Time `json:"updatedat"`
}

var Users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	Users = append(Users, newUser)
	json.NewEncoder(w).Encode(newUser)
}
