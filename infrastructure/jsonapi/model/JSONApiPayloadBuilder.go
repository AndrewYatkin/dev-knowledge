package jsonApiModel

import "dev-knowledge/infrastructure/errors"

const (
	LimitMetaKey      = "limit"
	OffsetMetaKey     = "offset"
	TotalItemsMetaKey = "totalItems"
)

type JSONApiPayloadBuilder struct {
	metaMap         map[string]interface{}
	dataObjectsList []*JSONApiObject
	dataObjectsMap  apiObjectsMap
	includeMap      map[ApiObjectKey]*JSONApiObject

	defaultsIncludes map[ApiObjectType]*JSONApiObject

	errors *errors.Errors
}

func NewJSONApiPayloadBuilder() *JSONApiPayloadBuilder {
	return &JSONApiPayloadBuilder{
		metaMap:          make(map[string]interface{}),
		dataObjectsList:  make([]*JSONApiObject, 0, 2),
		includeMap:       make(map[ApiObjectKey]*JSONApiObject),
		defaultsIncludes: make(map[ApiObjectType]*JSONApiObject),

		errors: errors.NewErrors(),
	}
}

func (b *JSONApiPayloadBuilder) AddPaginationMeta(meta *PaginationMeta) *JSONApiPayloadBuilder {
	if meta == nil {
		return b
	}

	b.AddMeta(LimitMetaKey, meta.Limit)
	b.AddMeta(OffsetMetaKey, meta.Offset)
	b.AddMeta(TotalItemsMetaKey, meta.TotalItems)

	return b
}

func (b *JSONApiPayloadBuilder) AddMeta(key string, value interface{}) *JSONApiPayloadBuilder {
	b.metaMap[key] = value

	return b
}

func (b *JSONApiPayloadBuilder) AddData(data ...*JSONApiObject) *JSONApiPayloadBuilder {
	for _, dataItem := range data {
		if dataItem == nil {
			continue
		}
		b.dataObjectsList = append(b.dataObjectsList, dataItem)
	}

	return b
}

func (b *JSONApiPayloadBuilder) AddInclude(includeData ...*JSONApiObject) *JSONApiPayloadBuilder {
	for _, includeDataItem := range includeData {
		if includeDataItem == nil {
			continue
		}
		objectKey := includeDataItem.Key()
		b.includeMap[objectKey] = includeDataItem
	}

	return b
}

func (b *JSONApiPayloadBuilder) AddDefaultsIncludes(objType string, defaults *JSONApiObject) *JSONApiPayloadBuilder {
	b.defaultsIncludes[ApiObjectType(objType)] = defaults
	return b
}

func (b *JSONApiPayloadBuilder) Build() (JSONApiPayload, error) {
	b.fillApiObjectsMap()
	b.fillNilRelationships()
	b.checkExistsIncludes()
	if b.errors.IsPresent() {
		return JSONApiPayload{}, b.errors
	}
	b.checkUsedIncludes()
	if b.errors.IsPresent() {
		return JSONApiPayload{}, b.errors
	}

	responsePayload := NewEmptyJSONApiPayload()
	b.setMetaToPayload(responsePayload)
	b.setDataToPayload(responsePayload)
	b.setIncludeToPayload(responsePayload)

	return *responsePayload, nil
}

func (b *JSONApiPayloadBuilder) fillApiObjectsMap() {
	b.dataObjectsMap = make(apiObjectsMap, 0, len(b.dataObjectsList))

	for _, dataItem := range b.dataObjectsList {
		objectKey := dataItem.Key()
		b.dataObjectsMap = b.dataObjectsMap.putObject(objectKey, dataItem)
	}
}

func (b *JSONApiPayloadBuilder) fillNilRelationships() {
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		if data.Relationships == nil {
			data.Relationships = NewEmptyJSONApiRelationships()
		}
	}

	for _, includeData := range b.includeMap {
		if includeData.Relationships == nil {
			includeData.Relationships = NewEmptyJSONApiRelationships()
		}
	}
}

func (b *JSONApiPayloadBuilder) checkExistsIncludes() {
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		b.checkExistsIncludesForRelationships(data.Relationships)
	}
	for _, include := range b.includeMap {
		b.checkExistsIncludesForRelationships(include.Relationships)
	}
}

func (b *JSONApiPayloadBuilder) checkExistsIncludesForRelationships(relationships JSONApiObjectRelationships) {
	relationshipApiObjectWithRequireInclude, err := getRelationshipApiObjectsWithRequireInclude(relationships)
	if err != nil {
		b.errors.AddError(err)
		return
	}

	for _, relationshipItem := range relationshipApiObjectWithRequireInclude {
		relationshipKey := relationshipItem.Key()

		if _, exists := b.includeMap[relationshipKey]; !exists {
			if _, defaultsExist := b.defaultsIncludes[relationshipKey.objectType]; defaultsExist {
				defaults := *b.defaultsIncludes[relationshipKey.objectType]
				defaults.ID = relationshipItem.ID
				b.includeMap[relationshipKey] = &defaults
			} else {
				b.errors.AddError(ErrIncludeForRelationshipNotFound(relationshipItem.ID, relationshipItem.Type))
				continue
			}
		}
	}
}

func (b *JSONApiPayloadBuilder) checkUsedIncludes() {
	relationshipApiObjectsSet, err := b.getRelationshipApiObjectsSet()
	if err != nil {
		b.errors.AddError(err)
		return
	}

	for includeKey := range b.includeMap {
		if _, exists := relationshipApiObjectsSet[includeKey]; !exists {
			b.errors.AddError(ErrNotUsedInclude(includeKey.ObjectID(), string(includeKey.objectType)))
		}
	}
}

func (b *JSONApiPayloadBuilder) getRelationshipApiObjectsSet() (map[ApiObjectKey]struct{}, error) {
	relationshipApiObjectsSet := make(map[ApiObjectKey]struct{})

	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		err := fillRelationshipApiObjectsSet(relationshipApiObjectsSet, data.Relationships)
		if err != nil {
			return nil, err
		}
	}

	for _, include := range b.includeMap {
		err := fillRelationshipApiObjectsSet(relationshipApiObjectsSet, include.Relationships)
		if err != nil {
			return nil, err
		}
	}

	return relationshipApiObjectsSet, nil
}

func fillRelationshipApiObjectsSet(
	relationshipObjectsSet map[ApiObjectKey]struct{},
	relationships JSONApiObjectRelationships,
) error {
	relationshipApiObjects, err := getAllRelationshipApiObjects(relationships)
	if err != nil {
		return err
	}

	for _, relationshipApiObject := range relationshipApiObjects {
		relationshipKey := relationshipApiObject.Key()
		relationshipObjectsSet[relationshipKey] = struct{}{}
	}
	return nil
}

func (b *JSONApiPayloadBuilder) setMetaToPayload(responsePayload *JSONApiPayload) {
	for key, value := range b.metaMap {
		responsePayload.AddMeta(key, value)
	}
}

func (b *JSONApiPayloadBuilder) setDataToPayload(responsePayload *JSONApiPayload) {
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		responsePayload.AddData(data)
	}
}

func (b *JSONApiPayloadBuilder) setIncludeToPayload(responsePayload *JSONApiPayload) {
	for _, include := range b.includeMap {
		responsePayload.AddInclude(include)
	}
}

func getRelationshipApiObjectsWithRequireInclude(relationships JSONApiObjectRelationships) ([]JSONApiBaseObject, error) {
	relationshipsDataList := make([]JSONApiBaseObject, 0)

	for _, relationshipItem := range relationships {
		err := validateRelationship(relationshipItem)
		if err != nil {
			return nil, err
		}

		if relationshipItem.Data == nil {
			continue
		}
		if relationshipItem.Links != nil && relationshipItem.Links.Related != "" {
			continue
		}

		relationshipData, ok := relationshipItem.Data.(JSONApiBaseObject)
		if !ok {
			relationshipArrayData, ok := relationshipItem.Data.([]JSONApiBaseObject)
			if !ok {
				return nil, ErrUnsupportedRelationshipDataStruct(relationshipItem.Data)
			}

			relationshipsDataList = append(relationshipsDataList, relationshipArrayData...)
			continue
		}

		relationshipsDataList = append(relationshipsDataList, relationshipData)
	}

	return relationshipsDataList, nil
}

func getAllRelationshipApiObjects(relationships JSONApiObjectRelationships) ([]JSONApiBaseObject, error) {
	relationshipsDataList := make([]JSONApiBaseObject, 0)

	for _, relationshipItem := range relationships {
		err := validateRelationship(relationshipItem)
		if err != nil {
			return nil, err
		}

		if relationshipItem.Data == nil {
			continue
		}

		relationshipData, ok := relationshipItem.Data.(JSONApiBaseObject)
		if !ok {
			relationshipArrayData, ok := relationshipItem.Data.([]JSONApiBaseObject)
			if !ok {
				return nil, ErrUnsupportedRelationshipDataStruct(relationshipItem.Data)
			}

			relationshipsDataList = append(relationshipsDataList, relationshipArrayData...)
			continue
		}

		relationshipsDataList = append(relationshipsDataList, relationshipData)
	}

	return relationshipsDataList, nil
}

func validateRelationship(relationshipItem JSONApiObjectRelationship) error {
	if relationshipItem.Data != nil {
		return nil
	}

	if relationshipItem.Links == nil {
		return ErrRelationshipWithoutDataAndLinks
	}

	if relationshipItem.Links.Self == "" {
		return ErrRelationshipWithoutDataAndLinksSelf
	}

	return nil
}

type apiObjectsMap []apiObjectPair

func (m apiObjectsMap) putObject(key ApiObjectKey, object *JSONApiObject) apiObjectsMap {
	pair := apiObjectPair{
		key:    key,
		object: object,
	}

	i, exists := m.findObjectIndex(key)
	if exists {
		m[i] = pair
		return m
	} else {
		return append(m, pair)
	}
}

func (m apiObjectsMap) findObjectIndex(key ApiObjectKey) (index int, exists bool) {
	for i, apiObjectPairItem := range m {
		if apiObjectPairItem.key == key {
			return i, true
		}
	}
	return -1, false
}

type apiObjectPair struct {
	key    ApiObjectKey
	object *JSONApiObject
}
