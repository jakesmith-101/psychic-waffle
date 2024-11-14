package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	PostID          string    `json:"postID"`          //pk
	PostTitle       string    `json:"postTitle"`       //
	PostDescription string    `json:"postDescription"` //
	Votes           int32     `json:"votes"`           //
	AuthorID        string    `json:"authorID"`        // fk
	CreatedAt       time.Time `json:"createdAt"`       //
	UpdatedAt       time.Time `json:"updatedAt"`       //
}

func GetPosts(sortID bool) (*[]Post, error) {
	var posts []Post
	var sortType = "CreatedAt"
	if sortID {
		sortType = "Votes"
	}
	rows, err := PgxPool.Query(context.Background(), fmt.Sprintf("SELECT * FROM posts ORDER BY %s DESC LIMIT 20", sortType))
	if err != nil {
		return &posts, err
	}
	posts, err = pgx.CollectRows(rows, pgx.RowToStructByName[Post])
	return &posts, err
}
