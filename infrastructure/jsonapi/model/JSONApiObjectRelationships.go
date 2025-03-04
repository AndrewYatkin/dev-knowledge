package jsonApiModel

type JSONApiObjectRelationships map[string]JSONApiObjectRelationship

type JSONApiObjectRelationship struct {
	Links *JSONApiLinks `json:"links,omitempty"`
	Data  interface{}   `json:"data,omitempty"`
}

func NewEmptyJSONApiRelationships() JSONApiObjectRelationships {
	return make(JSONApiObjectRelationships, 0)
}

func (r JSONApiObjectRelationships) AddWithoutDataWithSelfLink(key string, selfLink string) {
	r[key] = JSONApiObjectRelationship{Links: &JSONApiLinks{
		Self: selfLink,
	}}
}

func (r JSONApiObjectRelationships) AddApiBaseObject(objectID, objectType string) {
	jsonApiBaseObject := NewJSONApiBaseObject(objectID, objectType)
	r[objectType] = JSONApiObjectRelationship{Data: jsonApiBaseObject}
}

func (r JSONApiObjectRelationships) AddApiBaseObjBySpecialKey(objectID, objectType, key string) {
	jsonApiBaseObject := NewJSONApiBaseObject(objectID, objectType)
	r[key] = JSONApiObjectRelationship{Data: jsonApiBaseObject}
}

func (r JSONApiObjectRelationships) AddApiBaseObjects(key string, objects []JSONApiBaseObject) {
	r[key] = JSONApiObjectRelationship{Data: objects}
}

func (r JSONApiObjectRelationships) AddRelationshipData(key string, data interface{}) {
	r[key] = JSONApiObjectRelationship{Data: data}
}

func (r JSONApiObjectRelationships) AddRelationshipDataWithRelatedLink(key string, data interface{}, relatedLink string) {
	r[key] = JSONApiObjectRelationship{
		Links: &JSONApiLinks{
			Related: relatedLink,
		},
		Data: data,
	}
}

func (r JSONApiObjectRelationships) AddApiBaseObjectWithRelatedLink(objectID, objectType string, relatedLink string) {
	jsonApiBaseObject := NewJSONApiBaseObject(objectID, objectType)
	r[objectType] = JSONApiObjectRelationship{
		Links: &JSONApiLinks{
			Related: relatedLink,
		},
		Data: jsonApiBaseObject,
	}
}
