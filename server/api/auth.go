package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
)

// SignupOutput represents the signup operation response.
type SignupOutput struct {
	Body struct {
		Token   string `json:"token" example:"jwt" doc:"Jwt token string for auth"`
		Message string `json:"message" example:"Hello, John!" doc:"Greeting message"`
	}
}

func Signup(api huma.API) {
	// Register POST /auth/signup
	huma.Register(api, huma.Operation{
		OperationID: "post-account",
		Method:      http.MethodPost,
		Path:        BuildPath("/auth/signup"),
		Summary:     "Create an account",
		Description: "Create an account by username and password",
		Tags:        []string{"Signup"},
	}, func(ctx context.Context, input *struct {
		Username string `path:"username" maxLength:"30" example:"John" doc:"Name of account"`
		Password string `path:"password" maxLength:"30" example:"pass123" doc:"Password of account"`
	}) (*SignupOutput, error) {
		resp := &SignupOutput{}
		hash, err := password.GenerateFromPassword(input.Password)
		if err != nil {
			return resp, err
		}
		userID, err := db.CreateUser(input.Username, hash)
		if err != nil {
			return resp, err
		}
		user, err := db.GetUser(userID)
		if err != nil {
			return resp, err
		}
		tokenString, err := CreateToken(*user)
		if err != nil {
			return resp, err
		}
		resp.Body.Token = tokenString
		resp.Body.Message = fmt.Sprintf("Hello, %s!", user.Nickname)
		return resp, err
	})
}

// LoginOutput represents the login operation response.
type LoginOutput struct {
	Body struct {
		Token   string `json:"token" example:"jwt" doc:"Jwt token string for auth"`
		Message string `json:"message" example:"Hello, John!" doc:"Greeting message"`
		UserID  string `json:"userID" example:"uuid" doc:"ID of user's account"`
	}
}

func Login(api huma.API) {
	// Register POST /auth/login
	huma.Register(api, huma.Operation{
		OperationID: "get-account",
		Method:      http.MethodPost,
		Path:        BuildPath("/auth/login"),
		Summary:     "Log into account",
		Description: "Log into account by username and password",
		Tags:        []string{"Login"},
	}, func(ctx context.Context, input *struct {
		Username string `path:"username" maxLength:"30" example:"John" doc:"Name of account"`
		Password string `path:"password" maxLength:"30" example:"pass123" doc:"Password of account"`
	}) (*LoginOutput, error) {
		resp := &LoginOutput{}
		user, err := db.GetUserByUsername(input.Username)
		if err != nil {
			return resp, err
		}
		match, err := password.ComparePasswordAndHash(input.Password, user.PasswordHash)
		if err != nil {
			return resp, err
		}
		if match {
			tokenString, err := CreateToken(*user)
			if err != nil {
				return resp, err
			}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", user.Nickname)
			resp.Body.UserID = user.UserID
			resp.Body.Token = tokenString
		}
		return resp, err
	})
}
