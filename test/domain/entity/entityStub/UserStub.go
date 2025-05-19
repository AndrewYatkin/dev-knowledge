package entityStub

import (
	agreementEntity "dev-knowledge/domain/entity/agreement"
	profileEntity "dev-knowledge/domain/entity/profile"
	"dev-knowledge/domain/entity/spec"
	userEntity "dev-knowledge/domain/entity/user"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

func GetUserID() *userEntity.UserID {
	return userEntity.NewUserID()
}

func GetEmptyUser() *userEntity.User {
	user, err := userEntity.CreateByRole(spec.UserRoles.Customer())
	if err != nil {
		panic(err)
	}

	return user
}

func GetEmptyUserBuilder() *userEntity.Builder {

	newAgreement, err := agreementEntity.NewAgreement()
	if err != nil {
		panic(err)
	}

	return userEntity.NewBuilder().
		Profile(profileEntity.NewEmptyProfile()).
		Role(spec.UserRoles.Customer()).
		Agreement(newAgreement)
}

func GetFullCompletedUser() *userEntity.User {
	user, err := GetFullCompletedUserBuilder().Build()
	if err != nil {
		panic(err)
	}

	return user
}

func GetFullCompletedUserBuilder() *userEntity.Builder {
	return userEntity.NewBuilder().
		Email(GetActivatedUserEmails()).
		Profile(GetFullProfile()).
		Role(GetUserRole()).
		Agreement(GetAgreement()).
		Password(GetPassword()).
		CreatedAt(commonTime.Now())
}
