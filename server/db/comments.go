package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Comment struct {
	CommentID   string    `json:"commentID"`   //pk
	CommentText string    `json:"commentText"` //
	Votes       int32     `json:"votes"`       // 32 bits, a little over 2 billion (unlikely to have more than this upvote a post)
	PostID      string    `json:"postID"`      // fk
	ParentID    string    `json:"parentID"`    // fk
	AuthorID    string    `json:"authorID"`    // fk
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

func GetComment(commentID string) (*Comment, error) {
	var comment Comment
	row, err := PgxPool.Query(context.Background(), "SELECT * FROM comments WHERE CommentID=$1;", commentID)
	if err != nil {
		return &comment, err
	}
	comment, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Comment])
	return &comment, err
}

func CreateComment(postID string, parentID string, authorID string, text string) (string, error) {
	comment := Comment{
		CommentID:   uuid.NewString(),
		PostID:      postID,
		ParentID:    parentID,
		CommentText: text,
		AuthorID:    authorID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Votes:       0,
	}

	query := `INSERT INTO comments (CommentID, CommentText, Votes, ParentID, PostID, AuthorID, CreatedAt, UpdatedAt) VALUES (@CommentID, @CommentText, @Votes, @ParentID, @PostID, @AuthorID, @CreatedAt, @UpdatedAt)`
	args := pgx.NamedArgs{
		"CommentID":   comment.CommentID,
		"CommentText": comment.CommentText,
		"Votes":       comment.Votes,
		"ParentID":    comment.ParentID,
		"PostID":      comment.PostID,
		"AuthorID":    comment.AuthorID,
		"CreatedAt":   comment.CreatedAt,
		"UpdatedAt":   comment.UpdatedAt,
	}
	_, err = PgxPool.Exec(context.Background(), query, args)
	return comment.CommentID, err
}
