package jsonApiModel

type ApiObjectsSet []*JSONApiObject

func ApiObjectsSetFromList(apiObjectsList []*JSONApiObject) ApiObjectsSet {
	apiObjects := make([]*JSONApiObject, len(apiObjectsList))
	copy(apiObjects, apiObjectsList)

	return apiObjects
}

func (s ApiObjectsSet) ContainsObj(key ApiObjectKey) bool {
	for _, apiObject := range s {
		if apiObject.Key() == key {
			return true
		}
	}

	return false
}

func (s ApiObjectsSet) GetObj(key ApiObjectKey) (*JSONApiObject, bool) {
	for _, apiObject := range s {
		if apiObject.Key() == key {
			return apiObject, true
		}
	}

	return nil, false
}
