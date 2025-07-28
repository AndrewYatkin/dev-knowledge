package entityStub

import (
	profileEntity2 "dev-knowledge/domain/entity/user/profile"
)

func GetFullProfile() *profileEntity2.Profile {
	return profileEntity2.NewBuilder().
		FirstName("FirstName").
		LastName("LastName").
		Patronymic("Patronymic").
		Build()
}
