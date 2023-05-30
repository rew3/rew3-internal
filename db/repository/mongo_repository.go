package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	service "github.com/rew3/rew3-base/app/service"
	rcUtil "github.com/rew3/rew3-base/pkg/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository[Entity any] struct {
	Collection        *mongo.Collection
	RepositoryContext *RepositoryContext
}

func (repo *MongoRepository[Entity]) jsonToBson(rawMessage json.RawMessage) (bson.M, error) {
	var bsonM bson.M
	err := bson.Unmarshal(rawMessage, &bsonM)
	if err != nil {
		log.Printf("Conversion failed: %v\n", err)
		return nil, errors.New("Invalid Json Value")
	}
	return bsonM, nil
}

func (repo *MongoRepository[Entity]) bsonToJson(bsonData bson.M) (json.RawMessage, error) {
	extendedJSON, err := bson.MarshalExtJSON(bsonData, true, false)
	if err != nil {
		return nil, err
	}
	var rawMessage json.RawMessage = extendedJSON
	return rawMessage, nil
}

func (repo *MongoRepository[Entity]) entityToBson(entity *Entity) (bson.M, error) {
	bsonData, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}
	var doc bson.M
	err = bson.Unmarshal(bsonData, &doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (repo *MongoRepository[Entity]) Insert(ctx context.Context, data *Entity) (*Entity, error) {
	log.Printf("Creating record...")
	doc, err := repo.entityToBson(data)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	ec, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("Request Context is not available")
	}
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteNewMeta(&doc, ctx)
	res, err := repo.Collection.InsertOne(ec, docWithMeta)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	log.Printf("Record created with id - %v\n", res.InsertedID.(primitive.ObjectID))

	var insertedDoc Entity
	err = repo.Collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&insertedDoc)
	if err != nil {
		// Handle the error
		return nil, err
	} else {
		return &insertedDoc, nil
	}
}

func (repo *MongoRepository[Entity]) Update(ctx context.Context, id string, data *Entity) (*Entity, error) {
	// TODO need to apply security filter.
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	doc, err := repo.entityToBson(data)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return nil, err
	}
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(&doc, ctx)
	var record Entity
	err = repo.Collection.FindOneAndUpdate(ctx, filter, docWithMeta, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Record not found with ID %s\n", id)
			return nil, errors.New("Record not found with ID" + id)
		}
		log.Printf("Failed to update record: %v\n", err)
		return nil, err
	}
	return &record, nil
}

func (repo *MongoRepository[Entity]) UpdateDataOnly(ctx context.Context, id string, data *Entity) (bool, error) {
	// TODO need to apply security filter.
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return false, err
	}
	filter := bson.M{"_id": objectID}
	doc, err := repo.entityToBson(data)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	_, err = repo.Collection.UpdateOne(ctx, filter, doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Record not found with ID %s\n", id)
			return false, errors.New("Record not found with ID" + id)
		}
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	return true, nil
}

func (repo *MongoRepository[Entity]) FindAndUpdate(ctx context.Context, selector json.RawMessage, data *Entity) (bool, error) {
	// TODO need to apply security filter.
	filter, err := repo.jsonToBson(selector)
	if(err != nil) {
		return false, err
	}
	doc, err := repo.entityToBson(data)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(&doc, ctx)
	var record Entity
	err = repo.Collection.FindOneAndUpdate(ctx, filter, docWithMeta, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Print("Record not found with provided selector")
			return false, errors.New("Record not found with provided selector")
		}
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	return true, nil
}

func (repo *MongoRepository[Entity]) UpdateWithRawData(ctx context.Context, selector json.RawMessage, data json.RawMessage) (bool, error) {
	// TODO need to apply security filter.
	filter, err := repo.jsonToBson(selector)
	if(err != nil) {
		return false, err
	}
	update, err := repo.jsonToBson(data)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	var record Entity
	err = repo.Collection.FindOneAndUpdate(ctx, filter, bson.M{"$set" : update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Print("Record not found with provided selector")
			return false, errors.New("Record not found with provided selector")
		}
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	return true, nil
}

func (repo *MongoRepository[Entity]) UnsetFields(ctx context.Context, selector json.RawMessage, fields []string, multipleDoc bool) (bool, error) {
	unset := bson.M{}
	for _, field := range fields {
		unset[field] = ""
	}
	update := bson.M{"$unset": unset}
	filter, err := repo.jsonToBson(selector)
	if err != nil {
		return false, err
	}
	if multipleDoc {
		res, err := repo.Collection.UpdateMany(ctx, filter, update)
		if err != nil {
			return false, err
		}
		log.Printf("Document fields unset - %v", res.MatchedCount)
		return true, nil
	} else {
		res, err := repo.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return false, err
		}
		log.Printf("Document fields unset - %v", res.MatchedCount)
		return true, nil
	}
}

func (repo *MongoRepository[Entity]) Archive(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return false, errors.New("provide document ID is not valid")
	}
	ec, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("Request Context is not available")
	}
	updateData := bson.M{
		"meta._archived_at": time.Now(),
		"meta._archived_by": ec.User.Id,
	}
	update := bson.M{"$set": updateData}
	res, err := repo.Collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return false, err
	}
	log.Printf("Document archived - %v", res.MatchedCount)
	return true, nil
}

func (repo *MongoRepository[Entity]) UnArchive(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return false, errors.New("provide document ID is not valid")
	}
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("Request Context is not available")
	}
	selector, err := repo.bsonToJson(bson.M{"_id": objectID})
	if err != nil {
		return false, errors.New("Unable to unArchive document")
	}
	fields := []string{"meta._archived_at", "meta._archived_by"}
	return repo.UnsetFields(ctx, selector, fields, false)
}

func (repo *MongoRepository[Entity]) Delete(ctx context.Context, id string) (bool, error) {
	// TODO need to apply security filter.
	objectID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		log.Printf("Invalid ID: %v\n", err1)
		return false, fmt.Errorf("Record not found with ID %s\n", id)
	}
	filter := bson.M{"_id": objectID}
	_, err := repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting contact: %v", err)
		return false, err
	}
	return true, nil
}

func (repo *MongoRepository[Entity]) FindAndDelete(ctx context.Context, selector json.RawMessage) (bool, error) {
	// TODO need to apply security filter.
	_, err := repo.Collection.DeleteOne(ctx, selector)
	if err != nil {
		log.Printf("Error deleting contact: %v", err)
		return false, err
	}
	return true, nil
}

func (repo *MongoRepository[Entity]) BulkInsert(ctx context.Context, data []*Entity) (bool, error) {
	return false, nil
}

func (repo *MongoRepository[Entity]) BulkUpdate(ctx context.Context, data []*Entity) (bool, error) {
	return false, nil
}

func (repo *MongoRepository[Entity]) FindById(ctx context.Context, id string) *Entity {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil
	}
	var document Entity
	repo.RepositoryContext.DataSecurity.
		err = repo.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Record not found with ID %s\n", id)
			return nil
		}
		log.Printf("Failed to get document: %v\n", err)
		return nil
	}
	return &document
}

func (repo *MongoRepository[Entity]) Find(ctx context.Context, param service.RequestParam) []*Entity {
	results := []*Entity{}
	cur, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error listing documents: %v", err)
		return results
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var document Entity
		err := cur.Decode(&document)
		if err != nil {
			log.Printf("Error decoding document: %v", err)
		}
		results = append(results, &document)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating over documents: %v", err)
		return nil
	}
	return results
}

func (repo *MongoRepository[Entity]) FindBySelector(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) []*Entity {
	results := []*Entity{}
	cur, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error listing documents: %v", err)
		return results
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var document Entity
		err := cur.Decode(&document)
		if err != nil {
			log.Printf("Error decoding document: %v", err)
		}
		results = append(results, &document)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating over documents: %v", err)
		return nil
	}
	return results
}

func (repo *MongoRepository[Entity]) Count(ctx context.Context, param service.RequestParam) int64 {

}

func (repo *MongoRepository[Entity]) CountBySelector(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) int64 {

}

func (repo *MongoRepository[Entity]) Aggregate(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) AggregateWithLookupJoin(ctx context.Context) []*Entity {

}

func (repo *MongoRepository[Entity]) FindByIdPublic(ctx context.Context, id string) *Entity {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil
	}
	var document Entity
	repo.RepositoryContext.DataSecurity.
		err = repo.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Record not found with ID %s\n", id)
			return nil
		}
		log.Printf("Failed to get document: %v\n", err)
		return nil
	}
	return &document
}

func (repo *MongoRepository[Entity]) FindPublic(ctx context.Context, param service.RequestParam) []*Entity {
	results := []*Entity{}
	cur, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error listing documents: %v", err)
		return results
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var document Entity
		err := cur.Decode(&document)
		if err != nil {
			log.Printf("Error decoding document: %v", err)
		}
		results = append(results, &document)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating over documents: %v", err)
		return nil
	}
	return results
}

func (repo *MongoRepository[Entity]) FindBySelectorPublic(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) []*Entity {
	results := []*Entity{}
	cur, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error listing documents: %v", err)
		return results
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var document Entity
		err := cur.Decode(&document)
		if err != nil {
			log.Printf("Error decoding document: %v", err)
		}
		results = append(results, &document)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Error iterating over documents: %v", err)
		return nil
	}
	return results
}

func (repo *MongoRepository[Entity]) CountPublic(ctx context.Context, param service.RequestParam) int64 {

}

func (repo *MongoRepository[Entity]) CountBySelectorPublic(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) int64 {

}