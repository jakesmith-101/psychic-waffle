package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
)

type UpdateUserOutput struct {
	Body struct {
		Message string `json:"message" example:"Successfully updated user!" doc:"Success message"`
	}
}

func UpdateUser(api huma.API) {
	// Register POST /user/update
	huma.Register(api, huma.Operation{
		OperationID: "update-user",
		Method:      http.MethodPost,
		Path:        BuildPath("/user/update"),
		Summary:     "Update a user account",
		Description: "Update a user account by user ID",
		Tags:        []string{"UpdateUser"},
	}, func(ctx context.Context, input *struct {
		Body struct {
			UserID       string `json:"userID" required:"true"`    // pk
			Nickname     string `json:"nickname" required:"false"` //
			PasswordHash string `json:"password" required:"false"` //
			RoleID       string `json:"roleID" required:"false"`   //
			Token        string `json:"token" required:"true"`     // jwt token
		}
	}) (*UpdateUserOutput, error) {
		fmt.Fprintf(os.Stderr, "Requested update user: %s\n", input.Body.UserID)
		resp := &UpdateUserOutput{}

		// FIXME: permissions check using token!!
		success, err := db.SetUser(db.UpdateUser{
			UserID:       input.Body.UserID,
			Nickname:     input.Body.Nickname,
			PasswordHash: input.Body.PasswordHash,
			RoleID:       input.Body.RoleID,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		var name string
		if input.Body.Nickname != "" {
			name = input.Body.Nickname
		} else {
			name = input.Body.UserID
		}
		if success {
			resp.Body.Message = fmt.Sprintf("Successfully updated user: %s", name)
		}

		return resp, err
	})
}
