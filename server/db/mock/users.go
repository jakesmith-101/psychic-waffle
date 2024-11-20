package mock

import (
	"context"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/util"
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
func MockAdmin() (string, error) {
	user, err := db.GetUserByUsername("admin123")
	userID := user.UserID
	if err != nil {
		var pass string
		pass, err = util.GenerateFromPassword("admin123")
		if err != nil {
			return userID, err // invalid user ID
		}
		userID, err = db.CreateUser("admin123", pass)
	}
	return userID, err
}
