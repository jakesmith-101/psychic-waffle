package mock

import (
	"context"

	"github.com/jakesmith-101/psychic-waffle/db"
)

// Depends upon Users table
func CreatePostTable() error {
	_, err := db.PgxPool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS posts (
			PostID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			PostTitle TEXT,
			PostDescription TEXT,
			Votes INTEGER,
			AuthorID UUID references users(UserID),
			CreatedAt DATE,
			UpdatedAt DATE
		);`,
	)
	return err
}
