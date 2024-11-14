package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
)

type GetPostsOutput struct {
	Body struct {
		Posts []db.Post `json:"posts"`
	}
}

func GetPosts(api huma.API) {
	// Register GET /posts
	CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts/{sortID}",
		Summary: "Get 20 latest posts",
	}, func(ctx context.Context, input *struct {
		SortID string `path:"sortID" required:"true"` //
	}) (*GetPostsOutput, error) {
		resp := &GetPostsOutput{}
		var posts *[]db.Post
		var err error
		if input.SortID == "latest" {
			posts, err = db.GetLatestPosts()
		} else if input.SortID == "popular" {
			posts, err = db.GetPopularPosts()
		} else {
			err = errors.New("type of 'type' is incorrect")
		}
		if err != nil {
			return resp, err
		}
		resp.Body.Posts = *posts
		return resp, nil
	})
}
