package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	UserID       string    `json:"id"`        // pk
	Username     string    `json:"username"`  // unique
	Nickname     string    `json:"nickname"`  //
	PasswordHash string    `json:"password"`  //
	RoleID       string    `json:"roleid"`    // fk
	CreatedAt    time.Time `json:"createdat"` //
	UpdatedAt    time.Time `json:"updatedat"` //
}

func GetUser(ID string) (*User, error) {
	var user User
	row, err := Conn.Query(context.Background(), "SELECT * FROM users WHERE UserID=$1;", ID)
	if err != nil {
		return &user, err
	}
	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	row, err := Conn.Query(context.Background(), "SELECT * FROM users WHERE Username=$1;", username)
	if err != nil {
		return &user, err
	}

	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

func CreateUser(username string, passwordHash string) (string, error) {
	role, err := GetRoleByName("User")
	if err != nil {
		return "", err
	}
	user := User{
		Username:     username,
		PasswordHash: passwordHash,
		Nickname:     username,
		RoleID:       role.RoleID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO users (Username, PasswordHash, Nickname, RoleID, AuthToken, CreatedAt, UpdatedAt) VALUES (@Username, @PasswordHash, @Nickname, @RoleID, @AuthToken, @CreatedAt, @UpdatedAt)`
	args := pgx.NamedArgs{
		"UserID":       user.UserID,
		"Username":     user.Username,
		"PasswordHash": user.PasswordHash,
		"Nickname":     user.Nickname,
		"RoleID":       user.RoleID,
		"CreatedAt":    user.CreatedAt,
		"UpdatedAt":    user.UpdatedAt,
	}
	_, err = Conn.Exec(context.Background(), query, args)
	return user.UserID, err
}

// TODO: setUser logic
func SetUser(user User) bool {
	return true
}
