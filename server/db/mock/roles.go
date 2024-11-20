package mock

import (
	"context"
	"fmt"
	"os"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/util/permissions"
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
	_, err := db.GetRoleByName("Guest")
	if err != nil {
		_, err = db.CreateRole("Guest", permissions.Guest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
	}

	_, err = db.GetRoleByName("User")
	if err != nil {
		_, err = db.CreateRole("User", permissions.User)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
	}

	_, err = db.GetRoleByName("Moderator")
	if err != nil {
		_, err = db.CreateRole("Moderator", 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
	}

	_, err = db.GetRoleByName("Administrator")
	if err != nil {
		_, err = db.CreateRole("Administrator", 0) // FIXME: no perms "invented" yet
	}

	return err
}
