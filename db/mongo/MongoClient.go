package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	db     *mongo.Database
	client *mongo.Client
}

/**
* Create new instance of MongoClient.
* It is advices to use single instance of MongoClient per application.
 */
func NewMongoClient(mongoURL string, mongoDatabase string) *MongoClient {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	db := client.Database(mongoDatabase)
	return &MongoClient{db: db, client: client}
}

/**
* Get MongoDB Collection instance for given collection name.
 */
func (client *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return client.db.Collection(collectionName)
}

/**
* Stop toe mongodb client.
* You may want to stop mongodb client instance when application stop.
 */
func (client *MongoClient) StopClient() {
	client.client.Disconnect(context.Background())
}
