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
	if err != nil {
		return err
	}
	err = CreatePostTable()
	if err != nil {
		return err
	}
	err = CreateCommentTable()
	return err
}
