package entityStub

import (
	"dev-knowledge/domain/entity/user/spec"
)

func GetUserRole() spec.UserRole {
	return spec.UserRoles.Customer()
}
