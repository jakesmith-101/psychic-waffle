package mock

import (
	"context"

	"github.com/jakesmith-101/psychic-waffle/db"
)

// Depends upon Posts table
func CreateCommentTable() error {
	_, err := db.PgxPool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS comments (
			CommentID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			CommentText TEXT,
			ParentID UUID references comments(CommentID),
			PostID UUID references posts(PostID),
			CreatedAt DATE,
			UpdatedAt DATE
		);`,
	)
	return err
}
