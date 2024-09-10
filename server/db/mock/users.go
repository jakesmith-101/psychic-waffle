package mock

import (
	"context"
	"fmt"
	"os"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
)

// Depends upon Roles table
func CreateUserTable() error {
	_, err := db.PgxPool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
			UserID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Username VARCHAR(50),
			Nickname VARCHAR(50),
			PasswordHash VARCHAR(128),
			RoleID UUID references roles(RoleID),
			CreatedAt DATE,
			UpdatedAt DATE,
			UNIQUE (Username)
		);`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	return err
}

// Depends upon mocked roles
func MockAdmin() error {
	pass, err := password.GenerateFromPassword("admin123")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	_, err = db.CreateUser("admin", pass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	return err
}
