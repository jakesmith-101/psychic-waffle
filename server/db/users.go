package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type User struct {
	UserID       string    `json:"userID"`    // pk
	Nickname     string    `json:"nickname"`  //
	PasswordHash string    `json:"password"`  //
	RoleID       string    `json:"roleID"`    // fk
	Username     string    `json:"username"`  // unique
	CreatedAt    time.Time `json:"createdAt"` //
	UpdatedAt    time.Time `json:"updatedAt"` //
}

func GetUser(ID string) (*User, error) {
	var user User
	row, err := PgxPool.Query(context.Background(), "SELECT * FROM users WHERE UserID=$1;", ID)
	if err != nil {
		return &user, err
	}
	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[User])
	return &user, err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	row, err := PgxPool.Query(context.Background(), "SELECT * FROM users WHERE Username=$1;", username)
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
		UserID:       uuid.NewString(),
		Username:     username,
		PasswordHash: passwordHash,
		Nickname:     username,
		RoleID:       role.RoleID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO users (UserID, Username, PasswordHash, Nickname, RoleID, CreatedAt, UpdatedAt) VALUES (@UserID, @Username, @PasswordHash, @Nickname, @RoleID, @CreatedAt, @UpdatedAt)`
	args := pgx.NamedArgs{
		"UserID":       user.UserID,
		"Username":     user.Username,
		"PasswordHash": user.PasswordHash,
		"Nickname":     user.Nickname,
		"RoleID":       user.RoleID,
		"CreatedAt":    user.CreatedAt,
		"UpdatedAt":    user.UpdatedAt,
	}
	_, err = PgxPool.Exec(context.Background(), query, args)
	return user.UserID, err
}

type UpdateUser struct {
	UserID       string `json:"id"`       // pk
	Nickname     string `json:"nickname"` //
	PasswordHash string `json:"password"` //
	RoleID       string `json:"roleid"`   // fk
}

func SetUser(user UpdateUser) (bool, error) {
	// check vars are set
	nickname := ""
	password := ""
	roleid := ""
	if user.Nickname != "" {
		nickname = "Nickname=@Nickname,"
	}
	if user.PasswordHash != "" {
		password = "PasswordHash=@PasswordHash,"
	}
	if user.RoleID != "" {
		roleid = "RoleID=@RoleID,"
	}

	// only update set vars
	query := fmt.Sprintf(
		`UPDATE users SET %s %s %s UpdatedAt=@UpdatedAt WHERE UserID=@UserID;`,
		nickname,
		password,
		roleid,
	)
	args := pgx.NamedArgs{
		"UserID":       user.UserID,
		"Nickname":     user.Nickname,
		"PasswordHash": user.PasswordHash,
		"RoleID":       user.RoleID,
		"UpdatedAt":    time.Now(),
	}
	cmd, err := PgxPool.Exec(context.Background(), query, args)
	rowsAff := cmd.RowsAffected()
	if rowsAff == 0 {
		if err != nil {
			return false, err
		} else {
			return false, errors.New("no rows affected")
		}
	} else {
		if err != nil || rowsAff == 1 {
			return true, err
		} else {
			return true, errors.New("multiple rows affected")
		}
	}
}
