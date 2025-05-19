package id

import (
	"dev-knowledge/common/domainPrimitive/primitive/generator"
	idPrimitive "dev-knowledge/common/domainPrimitive/primitive/id"
	"dev-knowledge/infrastructure/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EntityIDShould struct {
	suite.Suite
}

func TestEntityIDShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(EntityIDShould))
}

func (s *EntityIDShould) TestNewEntityID_ReturnID() {
	entityID := idPrimitive.NewEntityID()

	assert.NotEmpty(s.T(), entityID)
}

func (s *EntityIDShould) TestEntityIDFrom_EmptyStr_ReturnError() {
	entityID, err := idPrimitive.EntityIDFrom("")

	assert.Empty(s.T(), entityID)
	assert.Equal(s.T(), idPrimitive.ErrEntityIDIsEmpty, err)
}

func (s *EntityIDShould) TestEntityIDFrom_NotValidUUID_ReturnError() {
	testCases := []struct {
		notValidUUID string
	}{
		{notValidUUID: "fff20a04-dab8-4275-83d8-05840facb11"},
		{notValidUUID: "f48389c-5b2d-11ed-9b6a-0242ac120002"},
		{notValidUUID: "f483789c-5bd-11ed-9b6a-0242ac120002"},
		{notValidUUID: "f483789c-5b2d-11d-9b6a-0242ac120002"},
		{notValidUUID: "f483789c-5b2d-11ed-9ba-0242ac120002"},
		{notValidUUID: "123"},
		{notValidUUID: "abc"},
		{notValidUUID: "fff20a04-1234-4275-83d8-05840facb11"},
	}

	for _, testCase := range testCases {
		entityID, err := idPrimitive.EntityIDFrom(testCase.notValidUUID)

		assert.Empty(s.T(), entityID)
		assert.True(s.T(), errors.EqualByCode(err, idPrimitive.CreateEntityIDErrorCode))
	}
}

func (s *EntityIDShould) TestEntityIDFrom_ValidStr_ReturnID() {
	testCases := []struct {
		uuid string
	}{
		{uuid: generator.GenerateUUID()},
		{uuid: "fff20a04-dab8-4275-83d8-05840facb112"},
		{uuid: "f483789c-5b2d-11ed-9b6a-0242ac120002"},
	}

	for _, testCase := range testCases {
		entityID, err := idPrimitive.EntityIDFrom(testCase.uuid)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), testCase.uuid, entityID.String())
	}
}
