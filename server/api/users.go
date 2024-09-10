package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
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
			Nickname string `json:"nickname" required:"false"` //
			Password string `json:"password" required:"false"` //
			RoleID   string `json:"roleID" required:"false"`   //
			Token    string `json:"token" required:"true"`     // jwt token
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

		var newPass string
		if input.Body.Password != "" {
			newPass, err = password.GenerateFromPassword(input.Body.Password)
			if err != nil {
				return resp, err
			}
		}
		success, err := db.SetUser(db.UpdateUser{
			UserID:       userID,
			Nickname:     input.Body.Nickname,
			PasswordHash: newPass,
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
		} else {
			return resp, errors.New("failed to update user")
		}

		return resp, err
	})
}

type GetUserOutput struct {
	Body struct {
		UserID    string    `json:"userID"`    // pk
		Nickname  string    `json:"nickname"`  //
		RoleID    string    `json:"roleID"`    // fk
		Username  string    `json:"username"`  // unique
		CreatedAt time.Time `json:"createdAt"` //
		UpdatedAt time.Time `json:"updatedAt"` //
	}
}

func GetUser(api huma.API) {
	// Register POST /user/get
	huma.Register(api, huma.Operation{
		OperationID: "get-user",
		Method:      http.MethodGet,
		Path:        BuildPath("/user/get"),
		Summary:     "Get a user account",
		Description: "Get a user account by user ID",
		Tags:        []string{"GetUser"},
	}, func(ctx context.Context, input *struct {
		Body struct {
			UserID string `json:"userID" required:"true"`
		}
	}) (*GetUserOutput, error) {
		resp := &GetUserOutput{}
		user, err := db.GetUser(input.Body.UserID)
		if err != nil {
			return resp, err
		}
		resp.Body.UserID = user.UserID
		resp.Body.Username = user.Username
		resp.Body.Nickname = user.Nickname
		resp.Body.RoleID = user.RoleID
		resp.Body.CreatedAt = user.CreatedAt
		resp.Body.UpdatedAt = user.UpdatedAt
		return resp, nil
	})
}
