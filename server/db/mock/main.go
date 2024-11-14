package mock

func MockAll() error {
	var err error
	err = CreateRoleTable()
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
	if err != nil {
		return err
	}
	return nil
}
