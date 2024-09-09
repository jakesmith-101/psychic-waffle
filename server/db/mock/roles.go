package mock

import "github.com/jakesmith-101/psychic-waffle/db"

func CreateRoleTable() {
	// TODO: create role table if not exist
}

func MockRoles() {
	db.CreateRole("User", 0)
	db.CreateRole("Admin", 0) // FIXME: no perms "invented" yet
}
