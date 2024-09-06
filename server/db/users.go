package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type User struct {
	UserID       string    `json:"id"`        // pk
	Username     string    `json:"username"`  // unique
	Email        string    `json:"email"`     // unique
	Nickname     string    `json:"nickname"`  //
	PasswordHash string    `json:"password"`  //
	RoleID       string    `json:"roleid"`    // fk
	AuthToken    string    `json:"authtoken"` // unique?
	CreatedAt    time.Time `json:"createdat"` //
	UpdatedAt    time.Time `json:"updatedat"` //
}

func GetUser(ID string) (*User, error) {
	var user User
	row, err := Conn.Query(context.Background(), "SELECT * FROM USERS WHERE UserID=$1;", ID)
	if err != nil {
		return &user, err
	}
	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	row, err := Conn.Query(context.Background(), "SELECT * FROM USERS WHERE Username=$1;", username)
	if err != nil {
		return &user, err
	}

	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

func CreateUser(username string, passwordHash string) (string, error) {
	user := User{
		UserID:       uuid.NewString(),
		Username:     username,
		PasswordHash: passwordHash,
		Nickname:     username,
		Email:        "",
		RoleID:       "",
		AuthToken:    "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO USERS (UserID, Username, PasswordHash, Nickname, Email, RoleID, AuthToken, CreatedAt, UpdatedAt) VALUES (@UserID, @Username, @PasswordHash, @Nickname, @Email, @RoleID, @AuthToken, @CreatedAt, @UpdatedAt)`
	args := pgx.NamedArgs{
		"UserID":       user.UserID,
		"Username":     user.Username,
		"PasswordHash": user.PasswordHash,
		"Nickname":     user.Nickname,
		"Email":        user.Email,
		"RoleID":       user.RoleID,
		"AuthToken":    user.AuthToken,
		"CreatedAt":    user.CreatedAt,
		"UpdatedAt":    user.UpdatedAt,
	}
	_, err := Conn.Exec(context.Background(), query, args)
	return user.UserID, err
}

// TODO:
func SetUser(user User) bool {
	return true
}
