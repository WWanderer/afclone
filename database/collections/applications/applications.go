package applications

import (
	"context"
	"log"
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
	res, err := applicationsCollection.InsertOne(ctx, application) // TODO handle error
	return res, err
}

func UpdateOrInsert(id uuid.UUID, update interface{}) (*mongo.UpdateResult, error) {
	applicationsCollection := dbconfig.Instance.DB.Collection("applications")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "ID", Value: id.String()}}

	result, err := applicationsCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println(err)
	}

	if result.MatchedCount != 0 {
		log.Println("matched and replaced an existing document")
	}
	if result.UpsertedCount != 0 {
		log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}

	return result, err

}

func FindById(id uuid.UUID, result interface{}) *mongo.SingleResult {
	applicationsCollection := dbconfig.Instance.DB.Collection("applications")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "ID", Value: id.String()}}

	return applicationsCollection.FindOne(ctx, filter)
}
