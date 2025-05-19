package userEntity

import idPrimitive "dev-knowledge/common/domainPrimitive/primitive/id"

type UserID struct {
	value idPrimitive.EntityID
}

func NewUserID() *UserID {
	return &UserID{value: idPrimitive.NewEntityID()}
}

func UserIDFrom(idStr string) (*UserID, error) {
	entityID, err := idPrimitive.EntityIDFrom(idStr)
	if err != nil {
		return nil, err
	}

	return &UserID{value: entityID}, nil
}

func (id *UserID) String() string {
	return string(id.value)
}
