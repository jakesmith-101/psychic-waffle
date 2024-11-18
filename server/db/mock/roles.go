package mock

import (
	"context"

	"github.com/jakesmith-101/psychic-waffle/db"
)

func CreateRoleTable() error {
	_, err := db.PgxPool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS roles (
			RoleID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Name VARCHAR(25),
			Permissions INT,
			UNIQUE (Name)
		);`,
	)
	return err
}

func MockRoles() error {
	_, err := db.GetRoleByName("User")
	if err != nil {
		_, err = db.CreateRole("User", 0)
		if err != nil {
			return err
		}
	}

	_, err = db.GetRoleByName("Moderator")
	if err != nil {
		_, err = db.CreateRole("Moderator", 0)
		if err != nil {
			return err
		}
	}

	_, err = db.GetRoleByName("Administrator")
	if err != nil {
		_, err = db.CreateRole("Administrator", 0) // FIXME: no perms "invented" yet
	}

	return err
}
