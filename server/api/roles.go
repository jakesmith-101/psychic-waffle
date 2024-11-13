package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jakesmith-101/psychic-waffle/db"
)

type GetRoleOutput struct {
	Body struct {
		RoleID      string `json:"roleID"`      // pk
		Permissions int    `json:"permissions"` //
		Name        string `json:"name"`        // unique
	}
}

func GetRole(api huma.API) {
	// Register GET /role
	CreateEndpoint(api, EndpointArgs{
		Method:  http.MethodGet,
		Path:    "/roles/{roleID}",
		Summary: "Get a role by role ID",
	}, func(ctx context.Context, input *struct {
		RoleID string `path:"roleID" required:"true"`
	}) (*GetRoleOutput, error) {
		resp := &GetRoleOutput{}
		role, err := db.GetRole(input.RoleID)
		if err != nil {
			return resp, err
		}
		resp.Body.RoleID = role.RoleID
		resp.Body.Permissions = role.Permissions
		resp.Body.Name = role.Name
		return resp, nil
	})
}
