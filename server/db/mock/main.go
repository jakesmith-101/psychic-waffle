package mock

func MockAll(mock bool) error {
	var err error
	var users []string

	// Role setup
	err = CreateRoleTable()
	if err != nil {
		return err
	}
	if mock {
		err = MockRoles()
		if err != nil {
			return err
		}
	}

	// User setup
	err = CreateUserTable()
	if err != nil {
		return err
	}
	if mock {
		var userID string
		userID, err = MockAdmin()
		if err != nil {
			return err
		}
		users = append(users, userID)
	}

	// Post setup
	err = CreatePostTable()
	if err != nil {
		return err
	}
	if mock {
		err = MockPosts(users)
		if err != nil {
			return err
		}
	}

	// Comment setup
	err = CreateCommentTable()
	if err != nil {
		return err
	}

	return nil
}
