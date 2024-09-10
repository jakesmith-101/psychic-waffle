package mock

func MockAll() error {
	err := CreateRoleTable()
	err2 := MockRoles()
	err3 := CreateUserTable()
	err4 := MockAdmin()
	if err != nil {
		return err
	} else if err2 != nil {
		return err2
	} else if err3 != nil {
		return err3
	} else if err4 != nil {
		return err4
	}
	return err
}
