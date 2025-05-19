package idPrimitive

import "dev-knowledge/common/domainPrimitive/primitive/generator"

type EntityID string

func NewEntityID() EntityID {
	return EntityID(generator.GenerateUUID())
}

func EntityIDFrom(strID string) (EntityID, error) {
	if strID == "" {
		return "", ErrEntityIDIsEmpty
	}

	id, err := generator.UUIDFrom(strID)
	if err != nil {
		return "", ErrCreateEntityID(strID)
	}

	return EntityID(id), nil
}
