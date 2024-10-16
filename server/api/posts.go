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
	// Register GET /posts
	CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts",
		Summary: "Get 20 latest posts",
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

type GetPopularPostsOutput struct {
	Body struct {
		Posts []db.Post `json:"posts"`
	}
}

func GetPopularPosts(api huma.API) {
	// Register GET /posts/popular
	CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts/popular",
		Summary: "Get 20 popular posts",
	}, func(ctx context.Context, input *struct{}) (*GetPopularPostsOutput, error) {
		resp := &GetPopularPostsOutput{}
		posts, err := db.GetPopularPosts()
		if err != nil {
			return resp, err
		}
		resp.Body.Posts = *posts
		return resp, nil
	})
}
