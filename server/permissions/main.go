package permissions

const (
	GetPost int64 = 1 << iota
	GetComment
	GetUser
	GetRole
	CreatePost
	CreateComment
	CreateUser
	CreateRole
	UpdatePost
	UpdateComment
	UpdateUser
	UpdateRole
	DeletePost
	DeleteComment
	DeleteUser
	DeleteRole
)

/*
&  AND
|  OR
^  XOR
&^ AND NOT
<< LEFT  SHIFT
>> RIGHT SHIFT
*/
