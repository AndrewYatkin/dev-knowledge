package stub

import (
	"dev-knowledge/boundary/dto"
	commonTesting "dev-knowledge/infrastructure/testing"
)

func GetRandomUserResponseDTO() *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		UserID:   commonTesting.RandomUUID(),
		Username: commonTesting.RandomDefaultStr(),
		Email:    commonTesting.RandomEmail(),
	}
}

func GetRandomCreateUserDTO() *dto.CreateUserDTO {
	return &dto.CreateUserDTO{
		Username: commonTesting.RandomDefaultStr(),
		Email:    commonTesting.RandomEmail(),
		Password: commonTesting.RandomDefaultStr(),
	}
}

func GetCreateUserDTOWithoutEmail() *dto.CreateUserDTO {
	return &dto.CreateUserDTO{
		Username: commonTesting.RandomDefaultStr(),
		Email:    "",
		Password: commonTesting.RandomDefaultStr(),
	}
}

func GetCrateUserDTOWithoutPassword() *dto.CreateUserDTO {
	return &dto.CreateUserDTO{
		Username: commonTesting.RandomDefaultStr(),
		Email:    commonTesting.RandomEmail(),
		Password: "",
	}
}
