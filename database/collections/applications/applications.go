package applications

import (
	"context"
	"time"

	"github.com/WWanderer/afclone/database/dbconfig"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(application interface{}) (*mongo.InsertOneResult, error) {
	applicationsCollection := dbconfig.Instance.DB.Collection("applications")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return applicationsCollection.InsertOne(ctx, application)
}

func UpdateOrInsert(id uuid.UUID, update interface{}) (*mongo.UpdateResult, error) {
	applicationsCollection := dbconfig.Instance.DB.Collection("applications")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "ID", Value: id.String()}}

	return applicationsCollection.UpdateOne(ctx, filter, update, opts)
}

func FindById(id uuid.UUID, result interface{}) error {
	applicationsCollection := dbconfig.Instance.DB.Collection("applications")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "ID", Value: id.String()}}
	return applicationsCollection.FindOne(ctx, filter).Decode(result)
}
