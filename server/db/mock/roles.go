package mock

import (
	"context"
	"fmt"
	"os"

	"github.com/jakesmith-101/psychic-waffle/db"
)

func CreateRoleTable() error {
	_, err := db.PgxPool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS roles (
			RoleID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Name VARCHAR(50),
			Permissions INT,
			UNIQUE (Name)
		);`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	return err
}

func MockRoles() error {
	_, err := db.CreateRole("User", 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	_, err = db.CreateRole("Admin", 0) // FIXME: no perms "invented" yet
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	return err
}
