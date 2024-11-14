package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Comment struct {
	CommentID   string    `json:"commentID"`   //pk
	CommentText string    `json:"commentText"` //
	Votes       int32     `json:"votes"`       //
	PostID      string    `json:"postID"`      // fk
	ParentID    string    `json:"parentID"`    // fk
	CreatedAt   time.Time `json:"createdAt"`   //
	UpdatedAt   time.Time `json:"updatedAt"`   //
}

func GetComments(postID string, sortID bool) (*[]Comment, error) {
	var comments []Comment
	var sortType = "CreatedAt"
	if sortID {
		sortType = "Votes"
	}
	rows, err := PgxPool.Query(context.Background(), fmt.Sprintf("SELECT * FROM comments WHERE PostID=$1 ORDER BY %s DESC LIMIT=20", sortType), postID)
	if err != nil {
		return &comments, err
	}
	comments, err = pgx.CollectRows(rows, pgx.RowToStructByName[Comment])
	return &comments, err
}
