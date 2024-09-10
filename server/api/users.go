package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
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
			UserID       string `json:"userID"`   // pk
			Nickname     string `json:"nickname"` //
			PasswordHash string `json:"password"` //
			RoleID       string `json:"roleID"`   // fk
			Token        string `json:"token"`    // jwt token
		}
	}) (*UpdateUserOutput, error) {
		fmt.Fprintf(os.Stderr, "Requested update user: %s\n", input.Body.UserID)
		resp := &UpdateUserOutput{}

		return resp, nil
	})
}
