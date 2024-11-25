package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/bbalet/stopwords"
	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func PostEndpoints(api huma.API) error {
	var err error

	err = GetPost(api)
	if err != nil {
		return err
	}

	err = GetPosts(api)
	if err != nil {
		return err
	}

	err = CreatePost(api)
	if err != nil {
		return err
	}

	return err
}

type GetPostOutput struct {
	Body struct {
		PostID          string    `json:"postID"`          // pk
		Slug            string    `json:"slug"`            // pk
		PostTitle       string    `json:"postTitle"`       //
		PostDescription string    `json:"postDescription"` //
		Votes           int32     `json:"votes"`           //
		AuthorID        string    `json:"authorID"`        // fk
		CreatedAt       time.Time `json:"createdAt"`       //
		UpdatedAt       time.Time `json:"updatedAt"`       //
	}
}

func GetPost(api huma.API) error {
	// Register GET /posts
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts/{slug}",
		Summary: "Get post",
	}, func(ctx context.Context, input *struct {
		Slug string `path:"slug" required:"true"` //
	}) (*GetPostOutput, error) {
		resp := &GetPostOutput{}
		post, err := db.GetPostBySlug(input.Slug)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		resp.Body = *post
		fmt.Fprintf(os.Stdout, "Get Post: %s", input.Slug)
		return resp, nil
	})
}

type GetPostsOutput struct {
	Body struct {
		Posts []db.Post `json:"posts"`
	}
}

func GetPosts(api huma.API) error {
	// Register GET /posts
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts",
		Summary: "Get 20 latest or most popular posts",
	}, func(ctx context.Context, input *struct {
		SortID bool `query:"sort" doc:"Order by votes or by date"`
	}) (*GetPostsOutput, error) {
		resp := &GetPostsOutput{}
		posts, err := db.GetPosts(input.SortID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		resp.Body.Posts = *posts
		if input.SortID {
			fmt.Fprintf(os.Stdout, "Get Posts: Popular")
		} else {
			fmt.Fprintf(os.Stdout, "Get Posts: Latest")
		}
		return resp, nil
	})
}

type CreatePostOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

func CreatePost(api huma.API) error {
	var whitespace *regexp.Regexp
	var duplicate *regexp.Regexp
	var err error

	whitespace, err = regexp.Compile(`[^a-z0-9-_]+`) // matches all non-alphanumeric
	if err != nil {
		return err
	}
	duplicate, err = regexp.Compile(`--+`) // matches multiple consecutive hyphens (often created by removing stop words)
	if err != nil {
		return err
	}

	// Register POST /posts
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodPost,
		Path:    "/posts",
		Summary: "Create new post",
	}, func(ctx context.Context, input *struct {
		Body struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
	}) (*CreatePostOutput, error) {
		resp := &CreatePostOutput{}
		// unaccent title
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
		result, _, _ := transform.String(t, input.Body.Title)

		// remove unnecessary words
		reducedTitle := stopwords.CleanString(result, "en", true) // cleans html too but kinda irrelevant here?
		// hyphenate on every non-alphanumeric
		hyphenatedTitle := whitespace.ReplaceAllString(strings.ToLower(reducedTitle), "-")
		// remove unnecessary hyphens
		dedupedTitle := duplicate.ReplaceAllString(strings.Trim(hyphenatedTitle, "-"), "-")

		// db trigger checks slug for uniqueness and appends a numeric suffix
		db.CreatePost(dedupedTitle, input.Body.Title, input.Body.Description, "")
		return resp, nil
	})
}
