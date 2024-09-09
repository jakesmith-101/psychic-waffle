package mock

import (
	"context"
	"log"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
)

// Depends upon Roles table
func CreateUserTable() error {
	_, err := db.Conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
			UserID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Email VARCHAR(50),
			Nickname VARCHAR(50),
			PasswordHash VARCHAR(128),
			RoleID UUID references roles(RoleID),
			CreatedAt DATE,
			UpdatedAt DATE,
			UNIQUE (Email)
		);`,
	)
	return err
}

// Depends upon mocked roles
func MockAdmin() error {
	pass, err := password.GenerateFromPassword("admin")
	if err != nil {
		log.Fatal("admin: password error")
	}
	_, err = db.CreateUser("admin", pass)
	return err
}
