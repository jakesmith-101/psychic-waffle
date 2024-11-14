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
		PostID string `path:"postID" required:"true"` //
		SortID string `path:"sortID" required:"true"` //
	}) (*GetCommentsOutput, error) {
		resp := &GetCommentsOutput{}
		var comments *[]db.Comment
		var err error
		if input.SortID == "latest" {
			comments, err = db.GetLatestComments(input.PostID)
		} else if input.SortID == "popular" {
			comments, err = db.GetPopularComments(input.PostID)
		} else {
			err = errors.New("type of 'sortID' is incorrect")
		}
		if err != nil {
			return resp, err
		}
		resp.Body.Comments = *comments
		return resp, nil
	})
}
