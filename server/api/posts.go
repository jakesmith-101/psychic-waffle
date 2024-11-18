package api

import (
	"context"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func PostEndpoints(api huma.API) error {
	var err error = nil

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

type GetPostsOutput struct {
	Body struct {
		Posts []db.Post `json:"posts"`
	}
}

func GetPosts(api huma.API) error {
	// Register GET /posts
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts/{sortID}",
		Summary: "Get 20 latest posts",
	}, func(ctx context.Context, input *struct {
		SortID bool `path:"sortID" required:"true"` //
	}) (*GetPostsOutput, error) {
		resp := &GetPostsOutput{}
		var posts *[]db.Post
		var err error
		posts, err = db.GetPosts(input.SortID)
		if err != nil {
			return resp, err
		}
		resp.Body.Posts = *posts
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
	var reduce *regexp.Regexp
	var err error

	whitespace, err = regexp.Compile(`[^a-z0-9\-_]+`) // matches all non-alphanumeric
	if err != nil {
		return err
	}
	duplicate, err = regexp.Compile(`--+`) // matches multiple consecutive hyphens
	if err != nil {
		return err
	}
	reduce, err = regexp.Compile(``) // TODO: add regexp to select words to be removed for slug
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

		// hyphenate on every non-alphanumeric
		hyphenatedTitle := whitespace.ReplaceAllString(strings.ToLower(result), "-")
		// remove unnecessary hyphens
		reducedTitle := duplicate.ReplaceAllString(strings.Trim(hyphenatedTitle, "-"), "-")
		// remove unnecessary words
		slug := reduce.ReplaceAllString(reducedTitle, "")
		db.CreatePost(slug, input.Body.Title, input.Body.Description, "")
		return resp, nil
	})
}
