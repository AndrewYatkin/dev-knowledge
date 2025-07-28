package emailEntity

import idPrimitive "dev-knowledge/common/domainPrimitive/primitive/id"

type EmailID struct {
	value idPrimitive.EntityID
}

func NewEmailID() *EmailID {
	return &EmailID{value: idPrimitive.NewEntityID()}
}

func EmailIDFrom(idStr string) (*EmailID, error) {
	entityID, err := idPrimitive.EntityIDFrom(idStr)
	if err != nil {
		return nil, err
	}

	return &EmailID{value: entityID}, nil
}

func (id *EmailID) String() string {
	return string(id.value)
}
