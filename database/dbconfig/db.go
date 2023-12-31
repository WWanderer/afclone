package dbconfig

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var Instance MongoInstance

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO load from env or conf file
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@127.0.0.1:27017/?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB Connected!")

	Instance = MongoInstance{
		Client: client,
		DB:     client.Database("afclone"),
	}

}
