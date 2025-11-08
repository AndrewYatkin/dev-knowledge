package mongoMock

import (
	"context"
	commonMock "dev-knowledge/infrastructure/testing/mock"
)

type MongoRepositoryMock struct {
	*commonMock.BaseMock
}

func GetMongoRepository() *MongoRepositoryMock {
	return &MongoRepositoryMock{
		BaseMock: commonMock.NewBaseMock(),
	}
}

func (m *MongoRepositoryMock) Insert(ctx context.Context, collectionName string, data interface{}) (string, error) {
	result, err := m.ProcessMethod("Insert", collectionName, data)
	if err != nil {
		return "", err
	}

	if result != nil {
		id, ok := result.(string)
		if !ok {
			return "", commonMock.ErrCannotCastResult
		}

		return id, nil
	}

	return "", err
}
