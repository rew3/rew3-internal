package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	db            *mongo.Database
	dbName        string
	client        *mongo.Client
	clientOptions *options.ClientOptions
}

/**
* Create new instance of MongoClient.
* It is advices to use single instance of MongoClient per application.
 */
func NewMongoClient(mongoDatabase string, clientOptions *options.ClientOptions) *MongoClient {
	client := MongoClient{dbName: mongoDatabase, clientOptions: clientOptions}
	client.connect()
	return &client
}

/**
 * Connect to mongodb client.
 */
func (client *MongoClient) connect() {
	// Set up MongoDB connection
	// Set a timeout for the initial connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, client.clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		// Connection is not active or there was an error
		log.Println("ERROR: MongoDB Client connection is not connected.")
	} else {
		log.Println("SUCCESS: MongoDB Client is connected.")
	}
	client.db = mongoClient.Database(client.dbName)
	client.client = mongoClient
}

/**
* Get MongoDB Collection instance for given collection name.
 */
func (client *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	if client.client == nil || client.db == nil {
		log.Println("ERROR: MongoDB Client unavailable. Connecting...")
		client.connect()
	}
	return client.db.Collection(collectionName)
}

/**
* Stop toe mongodb client.
* You may want to stop mongodb client instance when application stop.
 */
func (client *MongoClient) StopClient() {
	client.client.Disconnect(context.Background())
}
