package userRepo

import (
	"context"
	userEntity "dev-knowledge/domain/entity/user"
	logInterface "dev-knowledge/infrastructure/logger/interface"
	mongoInterface "dev-knowledge/infrastructure/mongo/interface"
	mongoModel "dev-knowledge/infrastructure/mongo/model"
)

const (
	indexUserID  = "uniqUserId"
	indexUSerKey = "user_id"
)

type UserRepository struct {
	mongoRepo      mongoInterface.MongoRepository
	logPublisher   logInterface.LogPublisher
	collection     string
	errorProcessor *errorProcessor
}

func (r *UserRepository) InitRepository(ctx context.Context) error {
	index := &mongoModel.DBIndex{
		Collection: r.collection,
		Name:       indexUserID,
		Keys:       []string{indexUSerKey},
		Type:       mongoModel.DBIndexAsc,
		Uniq:       true,
	}
	return r.mongoRepo.TryCreate(ctx, index)
}

func (r *UserRepository) Update(ctx context.Context, user *userEntity.User) error {

}
