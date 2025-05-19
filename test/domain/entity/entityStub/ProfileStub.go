package entityStub

import profileEntity "dev-knowledge/domain/entity/profile"

func GetFullProfile() *profileEntity.Profile {
	return profileEntity.NewBuilder().
		FirstName("FirstName").
		LastName("LastName").
		Patronymic("Patronymic").
		Build()
}
