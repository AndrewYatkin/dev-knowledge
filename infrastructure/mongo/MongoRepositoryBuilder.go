package mongo

import (
	"dev-knowledge/infrastructure/errors"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepositoryBuilder struct {
	mongoDb        *mongo.Database
	traceEnable    *bool
	logPublisher   loggerInterface.LogPublisher
	standAloneMode bool

	errors *errors.Errors
}

func NewMongoRepositoryBuilder() *MongoRepositoryBuilder {
	return &MongoRepositoryBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *MongoRepositoryBuilder) MongoDb(mongoDb *mongo.Database) *MongoRepositoryBuilder {
	b.mongoDb = mongoDb
	return b
}

func (b *MongoRepositoryBuilder) StandAloneModeEnable(standAloneMode bool) *MongoRepositoryBuilder {
	b.standAloneMode = standAloneMode
	return b
}

func (b *MongoRepositoryBuilder) TraceEnable(traceEnable bool) *MongoRepositoryBuilder {
	b.traceEnable = &traceEnable
	return b
}

func (b *MongoRepositoryBuilder) LogPublisher(logPublisher loggerInterface.LogPublisher) *MongoRepositoryBuilder {
	b.logPublisher = logPublisher
	return b
}

func (b *MongoRepositoryBuilder) Build() (*MongoRepository, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *MongoRepositoryBuilder) checkRequiredFields() {
	if b.mongoDb == nil {
		b.errors.AddError(ErrMongoDBIsRequired)
	}
	if b.logPublisher == nil {
		b.errors.AddError(ErrLogPublisherIsRequired)
	}
}

func (b *MongoRepositoryBuilder) createFromBuilder() *MongoRepository {
	return &MongoRepository{
		mongoDb:        b.mongoDb,
		standAloneMode: b.standAloneMode,
		traceEnable:    *b.traceEnable,
		logPublisher:   b.logPublisher,
	}
}
