package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitMongoDatabase(url, mongoDBName string) (*mongo.Database, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, errors.Wrap(err, "error creating and connecting mongo client")
	}

	pingTimeout := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), pingTimeout)
	defer cancelFunc()

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "mongo ping error")
	}

	return mongoClient.Database(mongoDBName), nil
}
