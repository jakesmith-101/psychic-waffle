package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/util"
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
	// Register GET /users
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
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error(), err)
		}
		resp.Body.UserID = user.UserID
		resp.Body.Username = user.Username
		resp.Body.Nickname = user.Nickname
		resp.Body.RoleID = user.RoleID
		resp.Body.CreatedAt = user.CreatedAt
		resp.Body.UpdatedAt = user.UpdatedAt
		util.Log(false, "Get User: %s", input.UserID)
		return resp, nil
	})
}

type UpdateUserOutput struct {
	Body struct {
		Message string `json:"message" example:"Successfully updated user!" doc:"Success message"`
	}
}

// TODO: change to allow singular changes etc, replace post with patch?
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
			Token    string `json:"token" required:"true"`     // FIXME: jwt token // Authorization Header instead
		}
	}) (*UpdateUserOutput, error) {
		resp := &UpdateUserOutput{}

		err := util.VerifyToken(input.Body.Token)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error(), err)
		}
		claims, err := util.ExtractClaims(input.Body.Token)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error(), err)
		}
		userID := fmt.Sprint(claims["UserID"])
		util.Log(false, "Requested update user: %s", userID)

		var newPass string
		if input.Body.Password != "" {
			newPass, err = util.GenerateFromPassword(input.Body.Password)
			if err != nil {
				util.LogError(err)
				return resp, huma.Error500InternalServerError(err.Error(), err)
			}
		}
		success, err := db.SetUser(db.UpdateUser{
			UserID:       userID,
			Nickname:     input.Body.Nickname,
			PasswordHash: newPass,
			RoleID:       input.Body.RoleID,
		})
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError("failed to update user")
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
			return resp, huma.Error500InternalServerError("failed to update user")
		}

		return resp, nil
	})
}
