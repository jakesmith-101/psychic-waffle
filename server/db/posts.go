package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	PostID          string    `json:"postID"`          //pk
	PostTitle       string    `json:"postTitle"`       //
	PostDescription string    `json:"postDescription"` //
	AuthorID        string    `json:"authorID"`        // fk
	CreatedAt       time.Time `json:"createdAt"`       //
	UpdatedAt       time.Time `json:"updatedAt"`       //
}

func GetLatestPosts() (*[]Post, error) {
	var posts []Post
	rows, err := PgxPool.Query(context.Background(), "SELECT * FROM posts ORDER BY CreatedAt DESC LIMIT=20")
	if err != nil {
		return &posts, err
	}
	posts, err = pgx.CollectRows(rows, pgx.RowToStructByName[Post])
	return &posts, err
}
