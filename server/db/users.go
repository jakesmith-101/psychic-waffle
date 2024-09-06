package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
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

func GetUser(ID string) (*User, error) {
	var user User
	row, err := Conn.Query(context.Background(), "SELECT * FROM USERS WHERE ID=$1;", ID)
	if err != nil {
		return &user, err
	}
	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	row, err := Conn.Query(context.Background(), "SELECT * FROM USERS WHERE username=$1;", username)
	if err != nil {
		return &user, err
	}

	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

/*
func CreateUser(username string, passwordHash string) bool {
	user := User{
		ID:           uuid.NewString(),
		Username:     username,
		PasswordHash: passwordHash,
		Nickname:     username,
		Email:        "",
		RoleID:       "",
		AuthToken:    "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return true
}

func SetUser(user User) bool {
	return true
}
*/
