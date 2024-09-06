package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Role struct {
	ID          string `json:"id"`
	Permissions int    `json:"permissions"`
	Name        string `json:"name"`
}

func GetRole(ID string) (*Role, error) {
	var role Role
	row, err := Conn.Query(context.Background(), "SELECT * FROM ROLES WHERE ID=$1;", ID)
	if err != nil {
		return &role, err
	}
	role, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Role])
	return &role, err
}

/*
func CreateRole(name string, perms int) bool {
	role := Role{
		ID:          uuid.NewString(),
		Permissions: perms,
		Name:        name,
	}
	return true
}
*/
