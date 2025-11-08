package userRepo

import (
	"dev-knowledge/infrastructure/errors"
	logInterface "dev-knowledge/infrastructure/logger/interface"
	mongoInterface "dev-knowledge/infrastructure/mongo/interface"
)

type UserRepositoryBuilder struct {
	userRepository *UserRepository
	errors         *errors.Errors
}

func NewUserRepositoryBuilder() *UserRepositoryBuilder {
	return &UserRepositoryBuilder{
		userRepository: &UserRepository{},
		errors:         errors.NewErrors(),
	}
}

func (b *UserRepositoryBuilder) MongoRepository(mongoRepo mongoInterface.MongoRepository) *UserRepositoryBuilder {
	b.userRepository.mongoRepo = mongoRepo
	return b
}

func (b *UserRepositoryBuilder) LogPublisher(logPublisher logInterface.LogPublisher) *UserRepositoryBuilder {
	b.userRepository.logPublisher = logPublisher
	return b
}

func (b *UserRepositoryBuilder) Collection(collection string) *UserRepositoryBuilder {
	b.userRepository.collection = collection
	return b
}

func (b *UserRepositoryBuilder) Build() (*UserRepository, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.userRepository.errorProcessor = &errorProcessor{
		logPublisher: b.userRepository.logPublisher,
	}

	return b.userRepository, nil
}

func (b *UserRepositoryBuilder) checkRequiredFields() {
	if b.userRepository.mongoRepo == nil {
		b.errors.AddError(ErrMongoRepoIsRequired)
	}
	if b.userRepository.logPublisher == nil {
		b.errors.AddError(ErrLogPublisherIsRequired)
	}
	if b.userRepository.collection == "" {
		b.errors.AddError(ErrCollectionIsRequired)
	}
}
