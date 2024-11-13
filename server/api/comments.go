package api

import (
	"context"
	"errors"
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
		Path:    "/posts/{postID}/comments/{type}",
		Summary: "Get 20 latest comments",
	}, func(ctx context.Context, input *struct {
		Body struct {
			PostID string `json:"postID" required:"true"` //
			Type   string `json:"type" required:"true"`   //
		}
	}) (*GetCommentsOutput, error) {
		resp := &GetCommentsOutput{}
		var comments *[]db.Comment
		var err error
		if input.Body.Type == "latest" {
			comments, err = db.GetLatestComments(input.Body.PostID)
		} else if input.Body.Type == "popular" {
			comments, err = db.GetPopularComments(input.Body.PostID)
		} else {
			err = errors.New("type of 'type' is incorrect")
		}
		if err != nil {
			return resp, err
		}
		resp.Body.Comments = *comments
		return resp, nil
	})
}
