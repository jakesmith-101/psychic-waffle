package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// SignupOutput represents the signup operation response.
type SignupOutput struct {
	Body struct {
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
		Username        string `path:"username" maxLength:"30" example:"John" doc:"Name of account"`
		Password        string `path:"password" maxLength:"30" example:"pass123" doc:"Password of account"`
		ConfirmPassword string `path:"confirmpassword" maxLength:"30" example:"pass123" doc:"Confirm Password of account"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Username) // TODO: signup logic
		return resp, nil
	})
}

// LoginOutput represents the login operation response.
type LoginOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, John!" doc:"Greeting message"`
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
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Username) // TODO: login logic
		return resp, nil
	})
}
