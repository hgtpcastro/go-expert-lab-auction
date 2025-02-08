package database

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(
	ctx context.Context,
	document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	if args.Get(0) != nil {
		return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCollection) UpdateByID(
	ctx context.Context,
	id interface{},
	update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, id, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) FindOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) Find(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}
