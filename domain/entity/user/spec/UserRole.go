package spec

type UserRole string
type UserRoleEnum map[string]UserRole

func (u UserRole) String() string {
	return string(u)
}

const (
	customerUserRole    = "customer"
	staffMemberUserRole = "staffMember"
)

var UserRoles = UserRoleEnum{
	customerUserRole:    customerUserRole,
	staffMemberUserRole: staffMemberUserRole,
}

func (e UserRoleEnum) Customer() UserRole {
	return e[customerUserRole]
}

func (e UserRoleEnum) StaffMember() UserRole {
	return e[staffMemberUserRole]
}

func (e UserRoleEnum) Of(code string) (UserRole, error) {
	role, ok := e[code]
	if !ok {
		return "", ErrUnsupportedUserRole(code)
	}

	return role, nil
}
