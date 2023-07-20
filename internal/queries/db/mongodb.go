package db

import (
	"RestApiStabDiffProject/internal/queries"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
}

func (d db) Create(ctx context.Context, query queries.Query) (string, error) {
	result, err := d.collection.InsertOne(ctx, query)
	if err != nil {
		return "", fmt.Errorf("Failed to insert a query due to an error %v", err)
	}
	oid, ok := result.InsertedID(primitive.ObjectID)
	if ok {
		return oid.Hex(), nill
	}
	return "", fmt.Errorf("failed to convert query to hex due to error: %v", err)
}

func (d db) Read(ctx context.Context, id string) (q queries.Query, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return q, fmt.Errorf("failed to convert hex to objectId, hex: %s", id)
	}
}

func (d db) Update(ctx context.Context, query queries.Query) error {
	//TODO implement me
	panic("implement me")
}

func (d db) Delete(ctx context.Context, is string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string) queries.Storage {
	return &db{
		collection: database.Collection(collection),
	}
}
