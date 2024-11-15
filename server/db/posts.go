package db

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Post struct {
	PostID          string    `json:"postID"`          // pk
	Slug            string    `json:"slug"`            // pk
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

func GetPostBySlug(slug string) (*Post, error) {
	var post Post
	row, err := PgxPool.Query(context.Background(), "SELECT * FROM posts WHERE Slug=$1;", slug)
	if err != nil {
		return &post, err
	}
	post, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Post])
	return &post, err
}

func GetPost(postID string) (*Post, error) {
	var post Post
	row, err := PgxPool.Query(context.Background(), "SELECT * FROM posts WHERE PostID=$1;", postID)
	if err != nil {
		return &post, err
	}
	post, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Post])
	return &post, err
}

var whitespace = regexp.MustCompile(`[^a-z0-9\-_]+`) // FIXME: assumes only usable characters are english alphanumeric
var duplicate = regexp.MustCompile(`^-+|-+$|--+`)    // matches leading or trailing hyphens, or multiple consecutive hyphens
var reduce = regexp.MustCompile(``)                  // TODO: add regexp to select words to be removed for slug

func CreatePost(title string, description string, author string) (string, error) {
	// unaccent title
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, title)

	// hyphenate on every non-alphanumeric
	title2 := whitespace.ReplaceAllString(strings.ToLower(result), "-")
	// remove unnecessary hyphens
	title3 := duplicate.ReplaceAllString(title2, "")
	// remove unnecessary words
	slug := reduce.ReplaceAllString(title3, "")

	post := Post{
		PostID:          uuid.NewString(),
		PostTitle:       title,
		PostDescription: description,
		AuthorID:        author,
		Slug:            slug, // FIXME: slug is not necessarily unique
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Votes:           0,
	}

	query := `INSERT INTO posts (PostID, Slug, PostTitle, PostDescription, Votes, AuthorID, CreatedAt, UpdatedAt) VALUES (@PostID, @Slug, @PostTitle, @PostDescription, @Votes, @AuthorID, @CreatedAt, @UpdatedAt)`
	args := pgx.NamedArgs{
		"PostID":          post.PostID,
		"Slug":            post.Slug,
		"PostTitle":       post.PostTitle,
		"PostDescription": post.PostDescription,
		"Votes":           post.Votes,
		"AuthorID":        post.AuthorID,
		"CreatedAt":       post.CreatedAt,
		"UpdatedAt":       post.UpdatedAt,
	}
	_, err = PgxPool.Exec(context.Background(), query, args)
	return post.PostID, err
}
