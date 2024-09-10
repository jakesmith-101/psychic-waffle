package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Comment struct {
	CommentID   string    `json:"commentID"`   //pk
	CommentText string    `json:"commentText"` //
	PostID      string    `json:"postID"`      // fk
	ParentID    string    `json:"parentID"`    // fk
	CreatedAt   time.Time `json:"createdAt"`   //
	UpdatedAt   time.Time `json:"updatedAt"`   //
}

func GetLatestComments() (*[]Comment, error) {
	var comments []Comment
	rows, err := PgxPool.Query(context.Background(), "SELECT * FROM comments ORDER BY CreatedAt DESC LIMIT=20")
	if err != nil {
		return &comments, err
	}
	comments, err = pgx.CollectRows(rows, pgx.RowToStructByName[Comment])
	return &comments, err
}
