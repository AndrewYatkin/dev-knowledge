package jsonApiModel

type JSONApiObject struct {
	ID            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    JSONApiAttributes          `json:"attributes"`
	Relationships JSONApiObjectRelationships `json:"relationships"`
}

func (o JSONApiObject) Key() ApiObjectKey {
	return ApiObjectKeyFrom(o.ID, ApiObjectType(o.Type))
}

func (o JSONApiObject) AddAttribute(key string, data interface{}) {
	o.Attributes[key] = data
}

func (o JSONApiObject) AddRelationship(objectID, objectType string) {
	o.Relationships.AddApiBaseObject(objectID, objectType)
}

func (o JSONApiObject) AddRelationshipBySpecialKey(objectID, objectType, key string) {
	o.Relationships.AddApiBaseObjBySpecialKey(objectID, objectType, key)
}

func (o JSONApiObject) GetRelationshipApiBaseObj(key string) (JSONApiBaseObject, error) {
	relationship, exists := o.Relationships[key]
	if !exists {
		return JSONApiBaseObject{}, ErrRelationshipByKeyNotFound(key, o.ID, o.Type)
	}

	apiBaseObj, isApiBaseObj := relationship.Data.(JSONApiBaseObject)
	if !isApiBaseObj {
		dataMap, isDataMap := relationship.Data.(map[string]interface{})
		if !isDataMap {
			return JSONApiBaseObject{}, ErrRelationshipDataIsNotApiBaseObj(key, o.ID, o.Type)
		}

		return o.convertMapToApiBaseObj(dataMap, key)
	}

	return apiBaseObj, nil
}

func (o JSONApiObject) GetRelationshipApiBaseObjects(key string) ([]JSONApiBaseObject, error) {
	relationship, exists := o.Relationships[key]
	if !exists {
		return nil, ErrRelationshipByKeyNotFound(key, o.ID, o.Type)
	}

	apiBaseObjects, isApiBaseObjects := relationship.Data.([]JSONApiBaseObject)
	if !isApiBaseObjects {
		return o.convertRelationshipToApiBaseObjects(relationship, key)
	}

	return apiBaseObjects, nil
}

func (o JSONApiObject) TryGetRelationshipApiBaseObjects(key string) ([]JSONApiBaseObject, error) {
	relationship, exists := o.Relationships[key]
	if !exists {
		return nil, ErrRelationshipByKeyNotFound(key, o.ID, o.Type)
	}

	apiBaseObjects, isApiBaseObjects := relationship.Data.([]JSONApiBaseObject)
	apiBaseObject, isApiBaseObject := relationship.Data.(JSONApiBaseObject)
	if !isApiBaseObjects && !isApiBaseObject {
		dataMap, isDataMap := relationship.Data.(map[string]interface{})
		if isDataMap {
			apiBaseObj, err := o.convertMapToApiBaseObj(dataMap, key)
			if err != nil {
				return nil, err
			}
			return []JSONApiBaseObject{apiBaseObj}, nil
		}

		return o.convertRelationshipToApiBaseObjects(relationship, key)
	}

	if isApiBaseObjects {
		return apiBaseObjects, nil
	} else {
		return []JSONApiBaseObject{apiBaseObject}, nil
	}
}

func (o JSONApiObject) convertRelationshipToApiBaseObjects(
	relationship JSONApiObjectRelationship,
	relationshipKey string,
) ([]JSONApiBaseObject, error) {
	apiBaseObjectMapList, isDataMap := relationship.Data.([]map[string]interface{})
	apiBaseObjectsList, isDataList := relationship.Data.([]interface{})
	if !isDataMap && !isDataList {
		return nil, ErrRelationshipDataIsNotApiBaseObjects(relationshipKey, o.ID, o.Type)
	}

	if isDataList {
		apiBaseObjectMapList = make([]map[string]interface{}, len(apiBaseObjectsList))
		for i, apiBaseObjectItem := range apiBaseObjectsList {
			apiBaseObjectMap, ok := apiBaseObjectItem.(map[string]interface{})
			if !ok {
				return nil, ErrRelationshipDataIsNotApiBaseObj(relationshipKey, o.ID, o.Type)
			}

			apiBaseObjectMapList[i] = apiBaseObjectMap
		}
	}

	apiBaseObjects := make([]JSONApiBaseObject, len(apiBaseObjectMapList))
	for i, apiBaseObjectMap := range apiBaseObjectMapList {
		apiBaseObj, err := o.convertMapToApiBaseObj(apiBaseObjectMap, relationshipKey)
		if err != nil {
			return nil, err
		}
		apiBaseObjects[i] = apiBaseObj
	}

	return apiBaseObjects, nil
}

func (o JSONApiObject) convertMapToApiBaseObj(dataMap map[string]interface{}, relationshipKey string) (JSONApiBaseObject, error) {
	objID, existsID := dataMap["id"]
	if !existsID {
		return JSONApiBaseObject{}, ErrRelationshipNotContainField(relationshipKey, "id", o.ID, o.Type)
	}
	objType, existsType := dataMap["type"]
	if !existsType {
		return JSONApiBaseObject{}, ErrRelationshipNotContainField(relationshipKey, "type", o.ID, o.Type)
	}

	objIDStr, ok := objID.(string)
	if !ok {
		return JSONApiBaseObject{}, ErrFailCastRelationshipField(relationshipKey, "id", o.ID, o.Type)
	}
	objTypeStr, ok := objType.(string)
	if !ok {
		return JSONApiBaseObject{}, ErrFailCastRelationshipField(relationshipKey, "type", o.ID, o.Type)
	}

	return JSONApiBaseObject{
		ID:   objIDStr,
		Type: objTypeStr,
	}, nil
}

func FindApiObjectFrom(apiObjects []*JSONApiObject, key ApiObjectKey) (*JSONApiObject, error) {
	for _, apiObject := range apiObjects {
		if apiObject.Key() == key {
			return apiObject, nil
		}
	}

	return &JSONApiObject{}, ErrApiObjectNotFound(key)
}
