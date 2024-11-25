package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
)

func CommentEndpoints(api huma.API) error {
	var err error = nil

	err = GetComments(api)
	if err != nil {
		return err
	}

	return err
}

type GetCommentsOutput struct {
	Body struct {
		Comments []db.Comment `json:"comments"`
	}
}

func GetComments(api huma.API) error {
	// Register GET /comments
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/posts/{postID}/comments/{sortID}",
		Summary: "Get 20 latest comments",
	}, func(ctx context.Context, input *struct {
		PostID string `path:"postID" required:"true"` //
		SortID bool   `path:"sortID" required:"true"` //
	}) (*GetCommentsOutput, error) {
		resp := &GetCommentsOutput{}
		var comments *[]db.Comment
		var err error
		comments, err = db.GetComments(input.PostID, input.SortID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		resp.Body.Comments = *comments
		if input.SortID {
			fmt.Fprintf(os.Stdout, "Get Comments: Popular")
		} else {
			fmt.Fprintf(os.Stdout, "Get Comments: Latest")
		}
		return resp, nil
	})
}
