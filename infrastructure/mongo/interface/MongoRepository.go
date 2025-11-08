package mongoInterface

import (
	"context"
	commonMongo "dev-knowledge/infrastructure/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionCallbackFunc func(ctx context.Context, repository MongoRepository) (interface{}, error)

type MongoRepository interface {
	Insert(ctx context.Context, collectionName string, data interface{}) (string, error)

	InsertMany(ctx context.Context, collectionName string, data []interface{}) ([]string, error)

	FindOneAndUpdate(
		ctx context.Context,
		collectionName string,
		resultModel,
		filter,
		updateData interface{},
		opt *options.FindOneAndUpdateOptions,
	) error

	ReplaceOne(
		ctx context.Context,
		collectionName string,
		filter,
		data interface{},
		opts ...*options.UpdateOptions,
	) (int64, error)

	UpdateOne(
		ctx context.Context,
		collectionName string,
		filter,
		data interface{},
		opts ...*options.UpdateOptions,
	) (int64, error)

	UpdateMany(
		ctx context.Context,
		collectionName string,
		filter interface{},
		data interface{},
		opt ...*options.UpdateOptions,
	) (int64, error)

	Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptions) error

	FindOne(
		ctx context.Context,
		collectionName string,
		resultModel,
		findQuery interface{},
		findOptions *options.FindOneOptions,
	) error

	DeleteOne(ctx context.Context, collectionName string, filter interface{},
		opt *options.DeleteOptions) (*mongo.DeleteResult, error)

	DeleteMany(ctx context.Context, collectionName string, filter interface{},
		opt *options.DeleteOptions) (*mongo.DeleteResult, error)

	Count(ctx context.Context, collectionName string, filter interface{}, opt *options.CountOptions) (int64, error)

	Aggregate(ctx context.Context, collectionName string, pipe mongo.Pipeline) (*mongo.Cursor, error)

	CreateIndex(ctx context.Context, index *commonMongo.DBIndex) error

	CollectionIndexes(ctx context.Context, collection string) (map[string]*commonMongo.DBIndex, error)
}
