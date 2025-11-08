package mongoDummy

import (
	"context"
	commonMongo "dev-knowledge/infrastructure/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositoryDummy struct{}

func GetMongoRepository() *MongoRepositoryDummy {
	return &MongoRepositoryDummy{}
}

func (m *MongoRepositoryDummy) Insert(ctx context.Context, collectionName string, data interface{}) (string, error) {
	return "", nil
}

func (m *MongoRepositoryDummy) InsertMany(ctx context.Context, collectionName string, data []interface{}) ([]string, error) {
	return nil, nil
}

func (m *MongoRepositoryDummy) FindOneAndUpdate(ctx context.Context, collectionName string, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptions) error {
	return nil
}

func (m *MongoRepositoryDummy) ReplaceOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.ReplaceOptions) error {
	return nil
}

func (m *MongoRepositoryDummy) UpdateOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.UpdateOptions) (int64, error) {
	return 0, nil
}

func (m *MongoRepositoryDummy) UpdateMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	return 0, nil
}

func (m *MongoRepositoryDummy) Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptions) error {
	return nil
}

func (m *MongoRepositoryDummy) FindOne(ctx context.Context, collectionName string, resultModel, findQuery interface{}, findOptions *options.FindOneOptions) error {
	return nil
}

func (m *MongoRepositoryDummy) DeleteOne(ctx context.Context, collectionName string, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	return nil, nil
}

func (m *MongoRepositoryDummy) DeleteMany(ctx context.Context, collectionName string, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	return nil, nil
}

func (m *MongoRepositoryDummy) Count(ctx context.Context, collectionName string, find interface{}, opt *options.CountOptions) (int64, error) {
	return 0, nil
}

func (m *MongoRepositoryDummy) Aggregate(ctx context.Context, collectionName string, pipe mongo.Pipeline) (*mongo.Cursor, error) {
	return nil, nil
}

func (m *MongoRepositoryDummy) CreateIndex(ctx context.Context, index *commonMongo.DBIndex) (string, error) {
	return "", nil
}

func (m *MongoRepositoryDummy) CollectionIndexes(ctx context.Context, collection string) (map[string]*commonMongo.DBIndex, error) {
	return nil, nil
}

func (m *MongoRepositoryDummy) TryCreateIndex(ctx context.Context, index *commonMongo.DBIndex) error {
	return nil
}
