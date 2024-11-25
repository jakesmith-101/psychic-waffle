package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/util"
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
	SetCookie [2]http.Cookie `header:"Set-Cookie"`
	Body      struct {
		Message string `json:"message" example:"Hello, John!" doc:"Greeting message"`
	}
}

func Signup(api huma.API) error {
	// Register POST /auth/signup
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodPost,
		Path:    "/auth/signup",
		Summary: "Create an account by username and password",
	}, func(ctx context.Context, input *AuthInput) (*AuthOutput, error) {
		util.Log(false, "Requested account creation: %s", input.Body.Username)
		resp := &AuthOutput{}
		hash, err := util.GenerateFromPassword(input.Body.Password)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		userID, err := db.CreateUser(input.Body.Username, hash)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		user, err := db.GetUser(userID)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		tokenString, err := util.CreateToken(*user)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", user.Nickname)
		resp.SetCookie = [2]http.Cookie{
			{
				Name:  "psychic_waffle_userid",
				Value: userID,
			},
			{
				Name:  "psychic_waffle_authorisation",
				Value: tokenString,
			},
		}
		return resp, nil
	})
}

func Login(api huma.API) error {
	// Register POST /auth/login
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodPost,
		Path:    "/auth/login",
		Summary: "Log into account by username and password",
	}, func(ctx context.Context, input *AuthInput) (*AuthOutput, error) {
		util.Log(false, "Requested account login: %s", input.Body.Username)
		resp := &AuthOutput{}
		user, err := db.GetUserByUsername(input.Body.Username)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		match, err := util.ComparePasswordAndHash(input.Body.Password, user.PasswordHash)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		if match {
			tokenString, err := util.CreateToken(*user)
			if err != nil {
				util.LogError(err)
				return resp, huma.Error500InternalServerError(err.Error())
			}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", user.Nickname)
			resp.SetCookie = [2]http.Cookie{
				{
					Name:  "psychic_waffle_userid",
					Value: user.UserID,
				},
				{
					Name:  "psychic_waffle_authorisation",
					Value: tokenString,
				},
			}
		}
		return resp, nil
	})
}
