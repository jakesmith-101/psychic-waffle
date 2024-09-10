package mock

import (
	"fmt"
	"os"
)

func MockAll() error {
	err := CreateRoleTable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	err = MockRoles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	err = CreateUserTable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	err = MockAdmin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
	return err
}
