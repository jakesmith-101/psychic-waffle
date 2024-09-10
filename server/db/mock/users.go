package mock

import (
	"context"

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
	return err
}

// Depends upon mocked roles
func MockAdmin() error {
	_, err := db.GetUserByUsername("admin123")
	if err != nil {
		var pass string
		pass, err = password.GenerateFromPassword("admin123")
		if err != nil {
			return err
		}
		_, err = db.CreateUser("admin123", pass)

	}
	return err
}
