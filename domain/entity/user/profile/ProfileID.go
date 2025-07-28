package profileEntity

import idPrimitive "dev-knowledge/common/domainPrimitive/primitive/id"

type ProfileID struct {
	value idPrimitive.EntityID
}

func NewProfileID() *ProfileID {
	return &ProfileID{value: idPrimitive.NewEntityID()}
}

func ProfileIDFrom(idStr string) (*ProfileID, error) {
	entityID, err := idPrimitive.EntityIDFrom(idStr)
	if err != nil {
		return nil, err
	}

	return &ProfileID{value: entityID}, nil
}

func (id *ProfileID) String() string {
	return string(id.value)
}
