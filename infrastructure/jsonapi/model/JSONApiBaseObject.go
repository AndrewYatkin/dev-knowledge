package jsonApiModel

type JSONApiBaseObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func NewJSONApiBaseObject(objectID, objectType string) JSONApiBaseObject {
	return JSONApiBaseObject{
		ID:   objectID,
		Type: objectType,
	}
}

func (o JSONApiBaseObject) Key() ApiObjectKey {
	return ApiObjectKeyFrom(o.ID, ApiObjectType(o.Type))
}
