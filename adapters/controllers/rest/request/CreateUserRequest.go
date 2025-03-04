package userRestRequest

import (
	"bytes"
	"dev-knowledge/boundary/dto"
	"encoding/json"
)

type CreateUserRequest struct {
	Data struct {
		Attributes struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"attributes"`
	} `json:"data"`
}

func (r *CreateUserRequest) FillFromBytes(jsonBytes []byte) error {
	return json.NewDecoder(bytes.NewReader(jsonBytes)).Decode(r)
}

func (r *CreateUserRequest) GetCreateUserDto() *dto.CreateUserDTO {
	return &dto.CreateUserDTO{
		Username: r.Data.Attributes.Username,
		Email:    r.Data.Attributes.Email,
		Password: r.Data.Attributes.Password,
	}
}
