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

func UserEndpoints(api huma.API) error {
	var err error = nil

	err = GetUser(api)
	if err != nil {
		return err
	}

	err = UpdateUser(api)
	if err != nil {
		return err
	}

	return err
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

func GetUser(api huma.API) error {
	// Register GET /user
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/users/{userID}",
		Summary: "Get a user account by user ID",
	}, func(ctx context.Context, input *struct {
		UserID string `path:"userID" required:"true"`
	}) (*GetUserOutput, error) {
		resp := &GetUserOutput{}
		user, err := db.GetUser(input.UserID)
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

type UpdateUserOutput struct {
	Body struct {
		Message string `json:"message" example:"Successfully updated user!" doc:"Success message"`
	}
}

func UpdateUser(api huma.API) error {
	// Register POST /user/update
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodPost,
		Path:    "/users/update",
		Summary: "Update a user account by user ID",
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
