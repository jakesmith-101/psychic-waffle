package mock

import (
	"github.com/ZeFort/chance"
	"github.com/jakesmith-101/psychic-waffle/util"
)

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
			util.LogError(err)
			return err
		}
	}

	// User setup
	err = CreateUserTable()
	if err != nil {
		return err
	}
	if mock {
		userID, err := MockAdmin()
		if err != nil {
			util.LogError(err)
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
			util.LogError(err)
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

var C *chance.Chance
var cIsSet = false
var cSeed int64

func NewChance(seed int64) {
	if !cIsSet || seed != cSeed {
		C = chance.NewS(seed)
		cIsSet = true
		cSeed = seed
	}
}
