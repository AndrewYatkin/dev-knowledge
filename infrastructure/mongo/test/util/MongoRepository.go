package mongoTestUtil

import (
	loggerFake "dev-knowledge/infrastructure/logger/test/fake"
	commonMongo "dev-knowledge/infrastructure/mongo"
	"fmt"
)

type MongoTestConfig struct {
	Host           string
	Port           string
	StandAloneMode bool
}

func (c *MongoTestConfig) MongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
}

func NewMongoRepository(config *MongoTestConfig, dbName string) *commonMongo.MongoRepository {
	mongoDB, err := commonMongo.InitMongoDatabase(config.MongoURI(), dbName)
	if err != nil {
		panic(err)
	}

	mongoRepository, err := commonMongo.NewMongoRepositoryBuilder().
		MongoDb(mongoDB).
		LogPublisher(loggerFake.GetLogPublisher()).
		StandAloneModeEnable(config.StandAloneMode).
		TraceEnable(false).
		Build()
	if err != nil {
		panic(err)
	}

	return mongoRepository
}
