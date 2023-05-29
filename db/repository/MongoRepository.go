package repository

import (
	"context"
	"fmt"
	"log"

	rcUtil "github.com/rew3/rew3-base/pkg/context"
	"github.com/rew3/rew3-base/app/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository[Entity any] struct {
	Collection        *mongo.Collection
	RepositoryContext *RepositoryContext
}

func (repo *MongoRepository[Entity]) Insert(ctx context.Context, data *Entity) (*Entity, error) {
	log.Printf("Creating record...")
	// Convert struct to BSON
	bsonData, err := bson.Marshal(data)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	var doc bson.M
	err = bson.Unmarshal(bsonData, &doc)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	ec := rcUtil.GetRequestContext(ctx)
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteNewMeta(&doc, ctx)
	res, err := repo.Collection.InsertOne(ec, docWithMeta)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	log.Printf("Record created with id - %v\n", res.InsertedID.(primitive.ObjectID))
	return data, nil
}

func (repo *MongoRepository[Entity]) Update(ctx context.Context, id string, data *Entity) (*Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	updateDoc := bson.M{}

	var record Entity
	err = repo.Collection.FindOneAndUpdate(ctx, filter, updateDoc, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Record not found with ID %s\n", id)
			return nil, fmt.Errorf("Record not found with ID %s\n", id)
		}
		log.Printf("Failed to update contact: %v\n", err)
		return nil, err
	}
	return &record, nil
}

func (repo *MongoRepository[Entity]) UpdateDataOnly(ctx context.Context, id string, data *Entity) (bool, error) {

}

func (repo *MongoRepository[Entity]) FindAndUpdate(ctx context.Context, selector bson.M, data *Entity) (bool, error) {

}

func (repo *MongoRepository[Entity]) UpdateWithRawData(ctx context.Context, selector bson.M, update bson.D) (bool, error) {

}

func (repo *MongoRepository[Entity]) UnsetFields(ctx context.Context, selector bson.M, fields []string, multipleDoc bool) (bool, error) {

}

func (repo *MongoRepository[Entity]) Archive(ctx context.Context, id string) (bool, error) {

}

func (repo *MongoRepository[Entity]) UnArchive(ctx context.Context, id string) (bool, error) {

}

func (repo *MongoRepository[Entity]) Delete(ctx context.Context, id string) (bool, error) {
	objectID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		log.Printf("Invalid ID: %v\n", err1)
		return false, fmt.Errorf("Record not found with ID %s\n", id)
	}
	filter := bson.M{"_id": objectID}
	_, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting contact: %v", err)
		return err
	}
	return nil
}

func (repo *MongoRepository[Entity]) FindAndDelete(ctx context.Context, selector bson.M) (bool, error) {

}

func (repo *MongoRepository[Entity]) BulkInsert(ctx context.Context, data []*Entity) (bool, error) {

}

func (repo *MongoRepository[Entity]) BulkUpdate(ctx context.Context, data []*Entity) (bool, error) {

}

func (repo *MongoRepository[Entity]) FindById(ctx context.Context) *Entity {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var contact Contact
	err = cs.collection.FindOne(ctx, filter).Decode(&contact)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Contact not found with ID %s\n", id)
			return nil, fmt.Errorf("contact not found")
		}
		log.Printf("Failed to get contact: %v\n", err)
		return nil, err
	}
	return &contact, nil
}

func (repo *MongoRepository[Entity]) Find(ctx context.Context) []*Entity {
	var results []*Contact
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error listing contacts: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var contact Contact
		err := cur.Decode(&contact)
		if err != nil {
			log.Printf("Error decoding contact: %v", err)
			return nil, err
		}
		results = append(results, &contact)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating over contacts: %v", err)
		return nil, err
	}
	return results, nil
}

func (repo *MongoRepository[Entity]) FindBySelector(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) Count(ctx context.Context) int64 {

}

func (repo *MongoRepository[Entity]) CountBySelector(ctx context.Context) int64 {

}

func (repo *MongoRepository[Entity]) Aggregate(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) AggregateWithLookupJoin(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) FindByIdPublic(ctx context.Context) *Entity {

}

func (repo *MongoRepository[Entity]) FindPublic(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) FindBySelectorPublic(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) CountPublic(ctx context.Context) int64 {

}

func (repo *MongoRepository[Entity]) CountBySelectorPublic(ctx context.Context) int64 {

}
