package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Role struct {
	RoleID      string `json:"id"`          // pk
	Permissions int    `json:"permissions"` //
	Name        string `json:"name"`        // unique
}

func GetRole(ID string) (*Role, error) {
	var role Role
	row, err := Conn.Query(context.Background(), "SELECT * FROM roles WHERE RoleID=$1;", ID)
	if err != nil {
		return &role, err
	}
	role, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Role])
	return &role, err
}

func CreateRole(name string, perms int) (string, error) {
	role := Role{
		RoleID:      uuid.NewString(),
		Permissions: perms,
		Name:        name,
	}

	query := `INSERT INTO users (RoleID, Permissions, Name) VALUES (@RoleID, @Permissions, @Name)`
	args := pgx.NamedArgs{
		"RoleID":      role.RoleID,
		"Permissions": role.Permissions,
		"Name":        role.Name,
	}
	_, err := Conn.Exec(context.Background(), query, args)
	return role.RoleID, err
}

func GetRoleByName(name string) (*Role, error) {
	var role Role
	row, err := Conn.Query(context.Background(), "SELECT * FROM roles WHERE Name=$1;", name)
	if err != nil {
		return &role, err
	}
	role, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Role])
	return &role, err
}
