package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
)

type GetCommentsOutput struct {
	Body struct {
		Comments []db.Comment `json:"comments"`
	}
}

func GetComments(api huma.API) {
	// Register GET /comments
	CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/comments",
		Summary: "Get 20 latest comments",
	}, func(ctx context.Context, input *struct {
		Body struct {
			PostID string `json:"postID" required:"true"` //
		}
	}) (*GetCommentsOutput, error) {
		resp := &GetCommentsOutput{}
		comments, err := db.GetLatestComments(input.Body.PostID)
		if err != nil {
			return resp, err
		}
		resp.Body.Comments = *comments
		return resp, nil
	})
}

type GetPopularCommentsOutput struct {
	Body struct {
		Comments []db.Comment `json:"comments"`
	}
}

func GetPopularComments(api huma.API) {
	// Register GET /comments/popular
	CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/comments/popular",
		Summary: "Get 20 popular comments",
	}, func(ctx context.Context, input *struct {
		Body struct {
			PostID string `json:"postID" required:"true"` //
		}
	}) (*GetPopularCommentsOutput, error) {
		resp := &GetPopularCommentsOutput{}
		comments, err := db.GetPopularComments(input.Body.PostID)
		if err != nil {
			return resp, err
		}
		resp.Body.Comments = *comments
		return resp, nil
	})
}
