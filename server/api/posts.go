package api

import (
	"context"
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
	// Register POST /user/get
	huma.Register(api, huma.Operation{
		OperationID: "get-posts",
		Method:      http.MethodGet,
		Path:        BuildPath("/posts"),
		Summary:     "Get 20 posts",
		Description: "Get 20 latest posts",
		Tags:        []string{"GetPosts"},
	}, func(ctx context.Context, input *struct{}) (*GetPostsOutput, error) {
		resp := &GetPostsOutput{}
		posts, err := db.GetLatestPosts()
		if err != nil {
			return resp, err
		}
		resp.Body.Posts = *posts
		return resp, nil
	})
}
