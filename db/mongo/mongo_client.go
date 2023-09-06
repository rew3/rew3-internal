package mongo

import (
	"context"
	"time"

	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
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
		logger.Log().Error("Failed to connect to MongoDB:", err)
	}
	client.db = mongoClient.Database(client.dbName)
	client.client = mongoClient
	if err := client.db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		logger.Log().Error("ERROR: MongoDB Client connection is not established: ", err)
	}
	logger.Log().Info("SUCCESS: MongoDB Client is connected.")
}

/**
* Get MongoDB Collection instance for given collection name.
 */
func (client *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	if client.client == nil || client.db == nil {
		logger.Log().Info("ERROR: MongoDB Client unavailable. Connecting...")
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
