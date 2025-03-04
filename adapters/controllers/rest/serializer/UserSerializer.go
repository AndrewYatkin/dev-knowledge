package serializer

import (
	"dev-knowledge/boundary/dto"
	jsonApiModel "dev-knowledge/infrastructure/jsonapi/model"
)

const (
	ResponseUser = "user"
)

func SerializeUser(user *dto.UserResponseDTO) (jsonApiModel.JSONApiPayload, error) {
	responseBuilder := jsonApiModel.NewJSONApiPayloadBuilder()
	responseBuilder.AddData(CreateUserObj(user))
	return responseBuilder.Build()
}

func CreateUserObj(user *dto.UserResponseDTO) *jsonApiModel.JSONApiObject {
	res := &jsonApiModel.JSONApiObject{
		ID:   user.UserID,
		Type: ResponseUser,
		Attributes: map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		},
		Relationships: jsonApiModel.JSONApiObjectRelationships{},
	}
	return res
}
