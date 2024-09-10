package mock

import (
	"fmt"
	"os"
)

func MockAll() error {
	err := CreateRoleTable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	err = MockRoles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	err = CreateUserTable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	err = MockAdmin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
