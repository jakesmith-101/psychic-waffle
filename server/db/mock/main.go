package mock

func MockAll() error {
	err := CreateRoleTable()
	if err != nil {
		return err
	}
	err = MockRoles()
	if err != nil {
		return err
	}
	err = CreateUserTable()
	if err != nil {
		return err
	}
	err = MockAdmin()
	return err
}
