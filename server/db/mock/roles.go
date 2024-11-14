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
			Name VARCHAR(50),
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
	_, err = db.GetRoleByName("Admin")
	if err != nil {
		_, err = db.CreateRole("Admin", 0) // FIXME: no perms "invented" yet
	}

	return err
}

/*
	TODO: Permissions list

	get post
	create post
	update post
	delete post

	get comment
	create comment
	update comment
	delete comment

	get user
	create user
	update user
	delete user

	get role
	create role
	update role
	delete role
*/
