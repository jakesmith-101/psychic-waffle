package mock

import (
	"context"

	"github.com/jakesmith-101/psychic-waffle/db"
)

func CreateRoleTable() error {
	_, err := db.Conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS roles (
			RoleID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Name VARCHAR(50),
			Permissions INT,
			UNIQUE (Name)
		);`,
	)
	return err
}

func MockRoles() error {
	_, err := db.CreateRole("User", 0)
	if err != nil {
		return err
	}
	_, err = db.CreateRole("Admin", 0) // FIXME: no perms "invented" yet
	return err
}
