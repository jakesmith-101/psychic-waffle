package main

type Permissions int64

const (
	GetPost Permissions = 1 << iota
	GetComment
	GetUser
	CreatePost
	CreateComment
	UpdatePost
	UpdateComment
	UpdateUser
	UpdateRole
	DeletePost
	DeleteComment
	DeleteUser
	DeleteRole
)

// FIXME: Permissions list might need reducing down to just get, create, update, delete
/*
	get post - Guest
	create post - User
	update post - Admin
	delete post - Admin

	get comment - Guest
	create comment - User
	update comment - Admin
	delete comment - Admin

	get user - Guest
	create user - Admin
	update user - Admin
	delete user - Admin

	get role - Admin
	create role - Admin
	update role - Admin
	delete role - Admin
*/
