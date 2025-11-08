package mongo

import (
	"context"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	mongoModel "dev-knowledge/infrastructure/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	indexTimeoutSeconds = 10
)

type handleOperationFunc func(ctx context.Context) (interface{}, error)

type MongoRepository struct {
	mongoDb        *mongo.Database
	standAloneMode bool
	traceEnable    bool
	logPublisher   loggerInterface.LogPublisher
}

func (r *MongoRepository) Insert(ctx context.Context, collectionName string, data interface{}) (string, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)

		res, err := coll.InsertMany(
			ctx,
			data)

		if err != nil {
			mongoErr, isMongoErr := err.(mongo.WriteException)
			if isMongoErr {
				for _, we := range mongoErr.WriteErrors {
					if r.isDuplicateError(we) {
						return nil, ErrDuplicateUniqueConstraint(err)
					}
					return nil, err
				}
			} else {
				return nil, err
			}
		}

		if res == nil || len(res.InsertIDs) == 0 {
			return nil, err
		}

		ids := make([]string, 0, len(res.InsertIDs))
		for _, id := range res.InsertIDs {
			if resID, ok := id.(primitive.ObjectID); ok {
				ids = append(ids, resID.Hex())
			}
		}

		return ids, nil
	}

	var id string
	if res != nil {
		if resID, ok := res.InsertID.(primitive.ObjectID); ok {
			id = resID.Hex()
		}
	}

	return id, nil
}

func (r *MongoRepository) InsertMany(ctx context.Context, collectionName string, data []interface{}) ([]string, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)

		res, err := coll.InsertMany(
			ctx,
			data)

		if err != nil {
			mongoErr, isMongoErr := err.(mongo.WriteException)
			if isMongoErr {
				for _, we := range mongoErr.WriteErrors {
					if r.isDuplicateError(we) {
						return nil, ErrDuplicateUniqueConstraint(err)
					}
					return nil, err
				}
			} else {
				return nil, err
			}
		}

		if res == nil || len(res, InsertIDs) == 0 {
			return nil, err
		}

		ids := make([]string, 0, len(res.InsertIDs))
		for _, id := range res.InsertIDs {
			if resID, ok := id.(primitive.ObjectID); ok {
				ids = append(ids, resID.Hex())
			}
		}

		return ids, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "InsertMany", handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.([]string), nil
}

func (r *MongoRepository) FindOneAndUpdate(
	ctx context.Context,
	collectionName string,
	resultModel,
	filter,
	updateData interface{},
	opt *options.FindOneAndUpdateOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
		result := coll.FindOneAndUpdate(
			ctx,
			filter,
			updateData,
			opt)

		err := result.Err()
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := r.handleOperation(ctx, collectionName, "FindOneAndUpdate", handleFunc)
	return err
}

func (r *MongoRepository) ReplaceOne(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
	opts ...*options.ReplaceOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
		_, err := coll.ReplaceOne(
			ctx,
			filter,
			data,
			opts...)
		return nil, err
	}

	_, err := r.handleOperation(ctx, collectionName, "ReplaceOne", handleFunc)
	return err
}

func (r *MongoRepository) UpdateOne(
	ctx context.Context,
	collectionName string,
	filter,
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
		res, err := coll.UpdateOne(
			ctx,
			filter,
			data,
			opts...)
		if err != nil {
			return nil, err
		}

		return res.ModifiedCount, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "UpdateOne", handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (r *MongoRepository) UpdateMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
		res, err := coll.UpdateMany(
			ctx,
			filter,
			data,
			opts...,
		)
		if err != nil {
			return 0, err
		}

		return res.ModifiedCount, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "UpdateMany", handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (r *MongoRepository) Find(
	ctx context.Context,
	collectionName string,
	results interface{},
	find interface{},
	opt *options.FindOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		cursor, err := collection.Find(ctx, find, opt)
		if err != nil {
			return nil, err
		}

		err = cursor.All(ctx, results)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := r.handleOperation(ctx, collectionName, "Find", handleFunc)
	return err
}

func (r *MongoRepository) FindOne(
	ctx context.Context,
	collectionName string,
	resultModel,
	findQuery interface{},
	findOptions *options.FindOneOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		result := collection.FindOne(ctx, findQuery, findOptions)
		err := result.Err()
		if err != nil {
			return nil, err
		}

		err = result.Decode(resultModel)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := r.handleOperation(ctx, collectionName, "FindOne", handleFunc)
	return err
}

func (r *MongoRepository) DeleteOne(
	ctx context.Context,
	collectionName string,
	filter interface{},
	opt *options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		result, err := collection.DeleteOne(ctx, filter, opt)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "DeleteOne", handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.DeleteResult), nil
}

func (r *MongoRepository) DeleteMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	opt *options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		result, err := collection.DeleteMany(ctx, filter, opt)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "DeleteMany", handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.DeleteResult), nil
}

func (r *MongoRepository) Count(
	ctx context.Context,
	collectionName string,
	find interface{},
	opt *options.CountOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		count, err := collection.CountDocuments(ctx, find, opt)
		if err != nil {
			return nil, err
		}

		return count, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "Count", handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (r *MongoRepository) Aggregate(
	ctx context.Context,
	collectionName string,
	pipe mongo.Pipeline,
) (*mongo.Cursor, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		cursor, err := collection.Aggregate(ctx, pipe)
		if err != nil {
			return nil, err
		}

		return cursor, nil
	}

	res, err := r.handleOperation(ctx, collectionName, "Aggregate", handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.Cursor), nil
}

func (r *MongoRepository) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	c := r.mongoDb.Collection(index.Collection)
	opts := options.CreateIndexes().SetMaxTime(indexTimeoutSeconds * time.Second)

	keysName := make([]bson.E, 0)
	for _, k := range index.Keys {
		keysName = append(keysName, bson.E{
			Key:   k,
			Value: int32(index.Type),
		})
	}
	keys := bson.D(keysName)
	indexModel := mongo.IndexModel{}
	indexModel.Keys = keys
	indexModel.Options = options.Index().SetName(index.Name)
	if index.Unique {
		indexModel.Options.SetUnique(true)
	}

	return c.Indexes().CreateOne(ctx, indexModel, opts)
}

func (r *MongoRepository) CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error) {
	res := make(map[string]*mongoModel.DBIndex)
	c := r.mongoDb.Collection(collection)
	duration := indexTimeoutSeconds * time.Second
	opts := &options.ListIndexesOptions{MaxTime: &duration}
	cur, err := c.Indexes().List(ctx, opts)
	if err != nil {
		return res, err
	}
	for cur.Next(ctx) {
		index := &mongoModel.DBIndex{}
		if err := cur.Decode(&index); err != nil {
			return res, err
		}
		res[index.Name] = index
	}
	return res, nil
}

func (r *MongoRepository) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	indexes, err := r.CollectionIndexes(ctx, index.Collection)
	if err != nil {
		return err
	}

	if r.isIndexExist(index, indexes) {
		return nil
	}

	_, err = r.CreateIndex(ctx, index)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoRepository) hadleOperation(ctx context.Context,
	collectionName string,
	methodName string,
	handleFunc handleOperationFunc,
) (interface{}, error) {
	res, err := handleFunc(ctx)
	if err != nil {
		handleErr := ErrHandleOperationInMongo(r.mongoDb.Name(), collectionName, methodName, err)
		r.logPublisher.LogError(ctx, handleErr)
	}

	return res, err
}

func (r *MongoRepository) isDuplicateError(we mongo.WriteError) bool {
	return we.Code == 11000
}

func (r *MongoRepository) isIndexExist(index *mongoModel.DBIndex, indexes map[string]*mongoModel.DBIndex) bool {
	_, ok := indexes[index.Name]
	return ok
}
