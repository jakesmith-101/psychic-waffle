package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func CreatePost(slug string, title string, description string, author string) (string, error) {
	post := Post{
		PostID:          uuid.NewString(),
		PostTitle:       title,
		PostDescription: description,
		AuthorID:        author,
		Slug:            slug,
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

func PostFuncs() error {
	var err error

	_, err = PgxPool.Exec(context.Background(), `
	CREATE OR REPLACE FUNCTION public.set_unique_slug() RETURNS trigger
		LANGUAGE plpgsql
		AS $$
	DECLARE
		base_slug TEXT;
		final_slug TEXT;
		counter INTEGER := 1;
	BEGIN
		base_slug := NEW.Slug;
		final_slug := base_slug;

		-- Loop to ensure uniqueness of the slug
		LOOP
			-- Check if the slug already exists in the table
			IF EXISTS (SELECT 1 FROM "posts" WHERE Slug = final_slug AND PostID != NEW.PostID) THEN
				-- If it exists, append a numeric suffix and increment the counter
				final_slug := base_slug || '-' || counter;
				counter := counter + 1;
			ELSE
				-- If it's unique, exit the loop
				EXIT;
			END IF;
		END LOOP;

		-- Set the unique slug to the 'Slug' field of the NEW record
		NEW.Slug := final_slug;
		RETURN NEW;
	END
	$$;
	`)
	if err != nil {
		return err
	}

	_, err = PgxPool.Exec(context.Background(), `
	CREATE OR REPLACE TRIGGER set_unique_slug
	BEFORE INSERT OR UPDATE
	ON "posts"
	FOR EACH ROW
	EXECUTE FUNCTION public.set_unique_slug();
	`)
	if err != nil {
		return err
	}

	return err
}
