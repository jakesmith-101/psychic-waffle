package mock

func MockAll() {
	CreateRoleTable()
	MockRoles()
	CreateUserTable()
	MockAdmin()
}
