package mock

import (
	"context"
	"fmt"
	"strings"

	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/util"
)

// Depends upon Users table
func CreatePostTable() error {
	_, err := db.PgxPool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS posts (
			PostID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			Slug TEXT,
			PostTitle TEXT,
			PostDescription TEXT,
			Votes INTEGER,
			AuthorID UUID references users(UserID),
			CreatedAt DATE,
			UpdatedAt DATE,
			UNIQUE (Slug)
		);`,
	)
	return err
}

func MockPosts(users []string) error {
	NewChance(0) // ensures C is set with seed
	length := C.IntBtw(20, 40)
	for i := 0; i < length; i++ {
		var slug string
		check := C.IntN(3)
		switch check {
		case 0:
			slug = C.Word()
		case 1:
			slug = fmt.Sprintf("%s-%s", C.Word(), C.Word())
		case 2:
			slug = fmt.Sprintf("%s-%s-%s", C.Word(), C.Word(), C.Word())
		}
		_, err := db.GetPostBySlug(slug)
		if err != nil {
			_, err = db.CreatePost(slug, strings.ReplaceAll(slug, "-", " "), C.String(), users[C.IntN(len(users))])
			if err != nil {
				return err
			} else {
				util.Log("ouput", "Mocked post: %s", slug)
			}
		}
	}
	return nil
}
