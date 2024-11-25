package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
	"github.com/jakesmith-101/psychic-waffle/util"
)

func RoleEndpoints(api huma.API) error {
	var err error = nil

	err = GetRole(api)
	if err != nil {
		return err
	}

	return err
}

type GetRoleOutput struct {
	Body struct {
		RoleID      string `json:"roleID"`      // pk
		Permissions int64  `json:"permissions"` //
		Name        string `json:"name"`        // unique
	}
}

func GetRole(api huma.API) error {
	// Register GET /role
	return CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/roles/{roleID}",
		Summary: "Get a role by role ID",
	}, func(ctx context.Context, input *struct {
		RoleID string `path:"roleID" required:"true"`
	}) (*GetRoleOutput, error) {
		resp := &GetRoleOutput{}
		role, err := db.GetRole(input.RoleID)
		if err != nil {
			util.LogError(err)
			return resp, huma.Error500InternalServerError(err.Error())
		}
		resp.Body.RoleID = role.RoleID
		resp.Body.Permissions = role.Permissions
		resp.Body.Name = role.Name
		util.Log(false, "Get Role: %s", input.RoleID)
		return resp, nil
	})
}
