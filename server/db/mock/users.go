package mock

import (
	"log"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
)

func CreateUserTable() {
	// TODO: create user table if not exist
}

func MockAdmin() {
	pass, err := password.GenerateFromPassword("admin")
	if err != nil {
		log.Fatal("admin: password error")
	}
	db.CreateUser("admin", pass)
}
