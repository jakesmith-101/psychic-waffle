package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/password"
)

func AuthEndpoints(api huma.API) error {
	var err error

	err = Signup(api)
	if err != nil {
		return err
	}

	err = Login(api)
	if err != nil {
		return err
	}

	return err
}

type AuthInput struct {
	Body struct {
		Username string `json:"username" example:"John" doc:"Name of account"`
		Password string `json:"password" example:"pass123" doc:"Password of account"`
	}
}

type AuthOutput struct {
	Body struct {
		Token   string `json:"token" example:"jwt" doc:"Jwt token string for auth"`
		Message string `json:"message" example:"Hello, John!" doc:"Greeting message"`
		UserID  string `json:"userID" example:"uuid" doc:"ID of user's account"`
	}
}

func Signup(api huma.API) error {
	// Register POST /auth/signup
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodPost,
		Path:    "/auth/signup",
		Summary: "Create an account by username and password",
	}, func(ctx context.Context, input *AuthInput) (*AuthOutput, error) {
		fmt.Fprintf(os.Stderr, "Requested account creation: %s\n", input.Body.Username)
		resp := &AuthOutput{}
		hash, err := password.GenerateFromPassword(input.Body.Password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		userID, err := db.CreateUser(input.Body.Username, hash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		user, err := db.GetUser(userID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		tokenString, err := CreateToken(*user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		resp.Body.Token = tokenString
		resp.Body.UserID = userID
		resp.Body.Message = fmt.Sprintf("Hello, %s!", user.Nickname)
		return resp, err
	})
}

func Login(api huma.API) error {
	// Register POST /auth/login
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodPost,
		Path:    "/auth/login",
		Summary: "Log into account by username and password",
	}, func(ctx context.Context, input *AuthInput) (*AuthOutput, error) {
		fmt.Fprintf(os.Stderr, "Requested account login: %s\n", input.Body.Username)
		resp := &AuthOutput{}
		user, err := db.GetUserByUsername(input.Body.Username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		match, err := password.ComparePasswordAndHash(input.Body.Password, user.PasswordHash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return resp, err
		}
		if match {
			tokenString, err := CreateToken(*user)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return resp, err
			}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", user.Nickname)
			resp.Body.UserID = user.UserID
			resp.Body.Token = tokenString
		}
		return resp, err
	})
}
