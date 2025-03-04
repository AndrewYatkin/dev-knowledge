package jsonApiModel

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
)

var (
	ErrRelationshipWithoutDataAndLinks     = errors.NewError("028be9b8-001", "Relationship data is nil and links is nil")
	ErrRelationshipWithoutDataAndLinksSelf = errors.NewError("028be9b8-002", "Relationship data is nil and links.self is empty")

	UnsupportedRelationshipDataStructErrorCode   errors.ErrorCode = "028be9b8-004"
	IncludeForRelationshipNotFoundErrorCode      errors.ErrorCode = "028be9b8-005"
	NotUsedIncludeErrorCode                      errors.ErrorCode = "028be9b8-006"
	RelationshipByKeyNotFoundErrorCode           errors.ErrorCode = "028be9b8-007"
	RelationshipDataIsNotApiBaseObjErrorCode     errors.ErrorCode = "028be9b8-008"
	RelationshipDataIsNotApiBaseObjectsErrorCode errors.ErrorCode = "028be9b8-009"
	ApiObjectNotFoundErrorCode                   errors.ErrorCode = "028be9b8-010"
	RelationshipNotContainFieldErrorCode         errors.ErrorCode = "028be9b8-011"
	FailCastRelationshipFieldErrorCode           errors.ErrorCode = "028be9b8-012"
)

func ErrUnsupportedRelationshipDataStruct(unsupportedData interface{}) error {
	errMsg := fmt.Sprintf("Unsupported relationship data struct. "+
		"Expected JSONApiBaseObject. Unsupported struct = '%s'", unsupportedData)
	return errors.NewError(UnsupportedRelationshipDataStructErrorCode, errMsg)
}

func ErrIncludeForRelationshipNotFound(objectID, objectType string) error {
	errMsg := fmt.Sprintf("Include for relationship not found. ID = '%s'. Type = '%s'", objectID, objectType)
	return errors.NewError(IncludeForRelationshipNotFoundErrorCode, errMsg)
}

func ErrNotUsedInclude(objectID, objectType string) error {
	errMsg := fmt.Sprintf("Include not used in relationships. ID = '%s'. Type = '%s'", objectID, objectType)
	return errors.NewError(NotUsedIncludeErrorCode, errMsg)
}

func ErrRelationshipByKeyNotFound(relationshipKey, objectID, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q not found in object with id = %q and type = %q",
		relationshipKey, objectID, objectType)
	return errors.NewError(RelationshipByKeyNotFoundErrorCode, errMsg)
}

func ErrRelationshipDataIsNotApiBaseObj(relationshipKey, objectID, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q is not JSONApiBaseObject in object with id = %q and type = %q",
		relationshipKey, objectID, objectType)
	return errors.NewError(RelationshipDataIsNotApiBaseObjErrorCode, errMsg)
}

func ErrRelationshipDataIsNotApiBaseObjects(relationshipKey, objectID, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q is not []JSONApiBaseObject in object with id = %q and type = %q",
		relationshipKey, objectID, objectType)
	return errors.NewError(RelationshipDataIsNotApiBaseObjectsErrorCode, errMsg)
}

func ErrApiObjectNotFound(key ApiObjectKey) error {
	errMsg := fmt.Sprintf("ApiObject not found by id = %q and type = %q", key.objectID, key.objectType)
	return errors.NewError(ApiObjectNotFoundErrorCode, errMsg)
}

func ErrRelationshipNotContainField(relationshipKey, field, objectID, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q not contain field %q in object with id = %q and type = %q",
		relationshipKey, field, objectID, objectType)
	return errors.NewError(RelationshipNotContainFieldErrorCode, errMsg)
}

func ErrFailCastRelationshipField(relationshipKey, field, objectID, objectType string) error {
	errMsg := fmt.Sprintf("Fail cast relationship field value. RelationshipKey = %q. Field = %q. "+
		"ObjectID = %q. ObjectType = %q",
		relationshipKey, field, objectID, objectType)
	return errors.NewError(FailCastRelationshipFieldErrorCode, errMsg)
}
