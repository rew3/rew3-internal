package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	mongoUtility "github.com/rew3/rew3-internal/db/utils"
	rcUtil "github.com/rew3/rew3-internal/pkg/context"
	service "github.com/rew3/rew3-internal/service/common/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * Mongo Repository.
 *
 * Note: For Entity, _id field should be string type not ObjectID type. It will be automatically converted to
 * Object ID and while persisting to database, and to Hex string while retrieving from database.
 */
type MongoRepository[Entity any] struct {
	Collection        *mongo.Collection
	RepositoryContext *RepositoryContext
}

func (repo *MongoRepository[Entity]) Insert(ctx context.Context, data *Entity) (*Entity, error) {
	log.Printf("Creating record...")
	doc, err := mongoUtility.EntityToBson(data, true)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	rc, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteNewMeta(&doc, &rc)
	res, err := repo.Collection.InsertOne(ctx, docWithMeta)
	if err != nil {
		log.Printf("Failed to create record: %v\n", err)
		return nil, err
	}
	log.Printf("Record created with id - %v\n", res.InsertedID.(primitive.ObjectID))

	var insertedDoc bson.M
	if err = repo.Collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&insertedDoc); err != nil {
		return nil, err
	} else {
		return mongoUtility.BsonToEntity[Entity](insertedDoc, true)
	}
}

func (repo *MongoRepository[Entity]) Update(ctx context.Context, id string, data *Entity) (*Entity, error) {
	// TODO need to apply security filter.
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil, err
	}
	rc, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	filter := bson.M{"_id": objectID}
	doc, err := mongoUtility.EntityToBson(data, false)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return nil, err
	}
	delete(doc, "_id") // todo later delete meta etc. TODO, get by ID, read meta and add meta. 
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(&doc, &rc)
	var record bson.M
	err = repo.Collection.FindOneAndUpdate(ctx, filter, docWithMeta, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Record not found with ID %s\n", id)
			return nil, errors.New("Record not found with ID" + id)
		}
		log.Printf("Failed to update record: %v\n", err)
		return nil, err
	}
	return mongoUtility.BsonToEntity[Entity](record, true)
}

func (repo *MongoRepository[Entity]) UpdateDataOnly(ctx context.Context, id string, data *Entity) (bool, error) {
	// TODO need to apply security filter.
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return false, err
	}
	filter := bson.M{"_id": objectID}
	doc, err := mongoUtility.EntityToBson(data, false)
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
	filter, err := mongoUtility.JsonToBson(selector)
	if err != nil {
		return false, err
	}
	rc, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	doc, err := mongoUtility.EntityToBson(data, false)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	delete(doc, "_id")
	docWithMeta := repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(&doc, &rc)
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
	filter, err := mongoUtility.JsonToBson(selector)
	if err != nil {
		return false, err
	}
	/*rc, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("Request Context is not available")
	}*/
	update, err := mongoUtility.JsonToBson(data)
	if err != nil {
		log.Printf("Failed to update record: %v\n", err)
		return false, err
	}
	var record Entity
	err = repo.Collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
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
	filter, err := mongoUtility.JsonToBson(selector)
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
	selector, err := mongoUtility.BsonToJson(bson.M{"_id": objectID})
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
	//repo.RepositoryContext.DataSecurity.
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
	return 0
}

func (repo *MongoRepository[Entity]) CountBySelector(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) int64 {
	return 0
}

func (repo *MongoRepository[Entity]) Aggregate(ctx context.Context) []*Entity {
	return nil
}

func (repo *MongoRepository[Entity]) AggregateWithLookupJoin(ctx context.Context) []*Entity {
	return nil
}

func (repo *MongoRepository[Entity]) FindByIdPublic(ctx context.Context, id string) *Entity {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		return nil
	}
	var document Entity
	//repo.RepositoryContext.DataSecurity.
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
	return 0
}

func (repo *MongoRepository[Entity]) CountBySelectorPublic(ctx context.Context, selector json.RawMessage, offset int, limit int, sort json.RawMessage) int64 {
	return 0
}
