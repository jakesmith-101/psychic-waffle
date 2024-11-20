package permissions

var Guest = GetComment |
	GetPost |
	GetUser

var User = Guest |
	CreateComment |
	CreatePost

var Moderator = User |
	GetRole |
	UpdatePost |
	UpdateUser |
	DeleteComment

var Administrator = Moderator |
	CreateUser |
	CreateRole |
	UpdateComment |
	UpdateRole |
	DeletePost |
	DeleteUser |
	DeleteRole

	// FIXME: Permissions list might need reducing down to just get, create, update, delete
	/*
		get post - Guest
		create post - User
		update post - Mod
		delete post - Admin

		get comment - Guest
		create comment - User
		update comment - Admin
		delete comment - Mod

		get user - Guest
		create user - Admin
		update user - Mod
		delete user - Admin

		get role - Mod
		create role - Admin
		update role - Admin
		delete role - Admin
	*/
