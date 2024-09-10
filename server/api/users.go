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
			Nickname     string `json:"nickname" required:"false"` //
			PasswordHash string `json:"password" required:"false"` //
			RoleID       string `json:"roleID" required:"false"`   //
			Token        string `json:"token" required:"true"`     // jwt token
		}
	}) (*UpdateUserOutput, error) {
		resp := &UpdateUserOutput{}

		err := VerifyToken(input.Body.Token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		claims, err := ExtractClaims(input.Body.Token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		userID := fmt.Sprint(claims["UserID"])
		fmt.Fprintf(os.Stderr, "Requested update user: %s\n", userID)

		success, err := db.SetUser(db.UpdateUser{
			UserID:       userID,
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
			name = userID
		}
		if success {
			resp.Body.Message = fmt.Sprintf("Successfully updated user: %s", name)
		}

		return resp, err
	})
}
