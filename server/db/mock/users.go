package mock

import (
	"context"
	"log"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
)

func CreateUserTable() error {
	_, err := db.Conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
			UserID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Username VARCHAR(50),
			Email VARCHAR(50),
			Nickname VARCHAR(50),
			PasswordHash VARCHAR(128),
			RoleID UUID references roles(RoleID),
			CreatedAt DATE,
			UpdatedAt DATE,
			UNIQUE (Username, Email)
		);`,
	)
	return err
}

/*
/UserID       string    `json:"id"`        // pk
/Username     string    `json:"username"`  // unique
/Email        string    `json:"email"`     // unique
/Nickname     string    `json:"nickname"`  //
/PasswordHash string    `json:"password"`  //
/RoleID       string    `json:"roleid"`    // fk
/CreatedAt    time.Time `json:"createdat"` //
/UpdatedAt    time.Time `json:"updatedat"` //
*/

func MockAdmin() error {
	pass, err := password.GenerateFromPassword("admin")
	if err != nil {
		log.Fatal("admin: password error")
	}
	_, err = db.CreateUser("admin", pass)
	return err
}
