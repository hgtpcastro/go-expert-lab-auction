package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionAPI interface {
	InsertOne(
		ctx context.Context,
		document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	UpdateByID(
		ctx context.Context,
		id interface{},
		update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)

	FindOne(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult

	Find(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}
