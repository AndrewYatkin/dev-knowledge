package jsonApiModel

type ApiObjectType string

func (t ApiObjectType) String() string {
	return string(t)
}

type ApiObjectKey struct {
	objectID   string
	objectType ApiObjectType
}

func ApiObjectKeyFrom(
	objectID string,
	objectType ApiObjectType,
) ApiObjectKey {
	return ApiObjectKey{
		objectID:   objectID,
		objectType: objectType,
	}
}

func (k ApiObjectKey) ObjectID() string {
	return k.objectID
}

func (k ApiObjectKey) ObjectType() ApiObjectType {
	return k.objectType
}
