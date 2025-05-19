package entityStub

import "dev-knowledge/domain/entity/spec"

func GetUserRole() spec.UserRole {
	return spec.UserRoles.Customer()
}
