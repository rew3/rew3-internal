package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	mongoUtility "github.com/rew3/rew3-internal/db/utils"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
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

/*
 * Insert record.
 */
func (repo *MongoRepository[Entity]) Insert(ctx context.Context, data *Entity) (*Entity, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (*Entity, error) {
		doc, err := mongoUtility.EntityToBsonD(data, false, true)
		if err != nil {
			logger.Log().Error("Invalid Input: Failed to create record:", err)
			return nil, err
		}
		doc = removeInternalFields(doc)
		id, doc := mongoUtility.WriteObjectID(doc)
		doc = repo.RepositoryContext.MetaDataWriter.WriteNewMeta(id.Hex(), doc, &rc)
		res, err := repo.Collection.InsertOne(ctx, doc)
		if err != nil {
			logger.Log().Error("Failed to create record:", err)
			return nil, err
		}
		var entity Entity
		selector := bson.M{"_id": res.InsertedID}
		if err = repo.Collection.FindOne(ctx, selector).Decode(&entity); err != nil {
			return nil, err
		}
		return &entity, nil
	})
}

/*
 * Update record.
 * Note: Include original meta info in Entity.
 */
func (repo *MongoRepository[Entity]) Update(ctx context.Context, id string, data *Entity) (*Entity, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (*Entity, error) {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Log().Error("Invalid Record ID:", err)
			return nil, err
		}
		doc, err := mongoUtility.EntityToBsonD(data, false, true)
		if err != nil {
			logger.Log().Error("Invalid Input: Failed to update record:", err)
			return nil, err
		}
		doc = removeInternalFields(doc)
		doc = repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(doc, &rc)
		update := bson.M{"$set": doc, "$inc": bson.M{"meta._version": 1}}
		var record Entity
		filter := bson.M{"_id": objectID}
		// TODO apply security.
		err = repo.Collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&record)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with ID", id)
				return nil, errors.New("Record not found with ID" + id)
			}
			logger.Log().Error("Failed to update record:", err)
			return nil, err
		}
		return &record, nil
	})
}

/*
 * Update record.
 * Note: Meta information is not updated.
 */
func (repo *MongoRepository[Entity]) UpdateOnly(ctx context.Context, id string, data *Entity) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Log().Error("Invalid Record ID:", err)
			return false, err
		}
		doc, err := mongoUtility.EntityToBsonD(data, false, true)
		if err != nil {
			logger.Log().Error("Invalid Input: Failed to update record:", err)
			return false, err
		}
		doc = removeInternalFields(doc)
		update := bson.M{"$set": doc}
		filter := bson.M{"_id": objectID}
		// TODO apply security.
		_, err = repo.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with ID:", id)
				return false, errors.New("Record not found with ID" + id)
			}
			logger.Log().Error("Failed to update record:", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) FindAndUpdate(ctx context.Context, selector bson.D, data *Entity, updateMany bool) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		doc, err := mongoUtility.EntityToBsonD(data, false, true)
		if err != nil {
			logger.Log().Error("Invalid Input: Failed to update record:", err)
			return false, err
		}
		doc = removeInternalFields(doc)
		doc = repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(doc, &rc)
		update := bson.M{"$set": doc, "$inc": bson.M{"meta._version": 1}}
		// TODO apply security.
		if updateMany {
			_, err = repo.Collection.UpdateMany(ctx, selector, update)
		} else {
			_, err = repo.Collection.UpdateOne(ctx, selector, update)
		}
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with provided selector")
				return false, errors.New("record not found with provided selector")
			}
			logger.Log().Error("Failed to update record:", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) FindAndUpdateBson(ctx context.Context, selector bson.D, data bson.D, updateMany bool) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		doc := removeInternalFields(data)
		doc = repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(doc, &rc)
		update := bson.M{"$set": doc, "$inc": bson.M{"meta._version": 1}}
		// TODO apply security.
		updateCall := func(isUpdateMany bool) error {
			if isUpdateMany {
				_, err := repo.Collection.UpdateMany(ctx, selector, update)
				return err
			} else {
				_, err := repo.Collection.UpdateOne(ctx, selector, update)
				return err
			}
		}
		err := updateCall(updateMany)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with provided selector")
				return false, errors.New("record not found with provided selector")
			}
			logger.Log().Error("Failed to update record:", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) UnsetFields(ctx context.Context, selector bson.D, fields []string, multipleDoc bool) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		// TODO - add security filter.
		unset := bson.M{}
		for _, field := range fields {
			unset[field] = ""
		}
		update := bson.M{"$unset": unset}
		if multipleDoc {
			res, err := repo.Collection.UpdateMany(ctx, selector, update)
			if err != nil {
				return false, err
			}
			logger.Log().Info("Document fields unset:", res.MatchedCount)
			return true, nil
		} else {
			res, err := repo.Collection.UpdateOne(ctx, selector, update)
			if err != nil {
				return false, err
			}
			logger.Log().Info("Document fields unset:", res.MatchedCount)
			return true, nil
		}
	})
}

func (repo *MongoRepository[Entity]) Archive(ctx context.Context, id string) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		// TODO - add security filter.
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Log().Error("Invalid ID:", err)
			return false, errors.New("provide document ID is not valid")
		}
		updateData := bson.M{
			"meta._archived_at": time.Now(),
			"meta._archived_by": rc.User.Id,
		}
		update := bson.M{"$set": updateData}
		res, err := repo.Collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
		if err != nil {
			return false, err
		}
		logger.Log().Info("Document archived:", res.MatchedCount)
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) UnArchive(ctx context.Context, id string) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		// TODO - add security filter.
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Log().Error("Invalid ID:", err)
			return false, errors.New("provide document ID is not valid")
		}
		selector := bson.D{{Key: "_id", Value: objectID}}
		fields := []string{"meta._archived_at", "meta._archived_by"}
		return repo.UnsetFields(ctx, selector, fields, false)
	})
}

func (repo *MongoRepository[Entity]) Delete(ctx context.Context, id string) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		// TODO - add security filter.
		objectID, err1 := primitive.ObjectIDFromHex(id)
		if err1 != nil {
			logger.Log().Error("Invalid ID:", err1)
			return false, fmt.Errorf("record not found with id %s", id)
		}
		filter := bson.M{"_id": objectID}
		_, err := repo.Collection.DeleteOne(ctx, filter)
		if err != nil {
			logger.Log().Error("Error deleting contact:v", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) FindAndDelete(ctx context.Context, selector bson.D, multipleDoc bool) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		// TODO - add security filter.
		if multipleDoc {
			if _, err := repo.Collection.DeleteMany(ctx, selector); err != nil {
				logger.Log().Error("Error deleting contact:", err)
				return false, err
			}
		} else {
			if _, err := repo.Collection.DeleteOne(ctx, selector); err != nil {
				logger.Log().Error("Error deleting contact:", err)
				return false, err
			}
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) BulkInsert(ctx context.Context, data []*Entity) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		bsonDocuments := []*bson.D{}
		for _, entity := range data {
			doc, err := mongoUtility.EntityToBsonD(&entity, false, true)
			if err != nil {
				return false, err
			}
			doc = removeInternalFields(doc)
			id, doc := mongoUtility.WriteObjectID(doc)
			doc = repo.RepositoryContext.MetaDataWriter.WriteNewMeta(id.Hex(), doc, &rc)
			bsonDocuments = append(bsonDocuments, &doc)
		}
		docs := make([]interface{}, len(bsonDocuments))
		for i, doc := range bsonDocuments {
			docs[i] = doc.Map()
		}
		_, err := repo.Collection.InsertMany(ctx, docs)
		if err != nil {
			logger.Log().Error("Failed to bulk create record:", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) BulkUpdate(ctx context.Context, data map[string]*Entity) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		var writes []mongo.WriteModel
		for key, entity := range data {
			filter := bson.M{"_id": key}
			doc, err := mongoUtility.EntityToBsonD(&entity, true, true)
			if err != nil {
				return false, fmt.Errorf("invalid entity data : %v", doc)
			}
			doc = removeInternalFields(doc)
			doc = repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(doc, &rc)
			update := bson.M{"$set": doc, "$inc": bson.M{"meta._version": 1}}
			updateModel := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
			writes = append(writes, updateModel)
		}
		_, err := repo.Collection.BulkWrite(context.TODO(), writes)
		if err != nil {
			logger.Log().Error("Failed to bulk create record:", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) BulkDelete(ctx context.Context, ids []string) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		// TODO - add security filter.
		// Convert the IDs to ObjectIDs
		objectIDs := make([]primitive.ObjectID, len(ids))
		for i, id := range ids {
			objectID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				logger.Log().Error("Unable to bulk delete: ", err)
			}
			objectIDs[i] = objectID
		}
		filter := bson.M{
			"_id": bson.M{
				"$in": objectIDs,
			},
		}
		_, err := repo.Collection.DeleteMany(ctx, filter)
		if err != nil {
			logger.Log().Error("Error bulk deleting records:", err)
			return false, err
		}
		return true, nil
	})
}

// elementUpdateValue must be in nested non flatten structure.
func (repo *MongoRepository[Entity]) AppendToArrayField(
	ctx context.Context,
	docId string,
	arrFieldPath string,
	elementSelectorName string,
	elementSelectorValue string,
	elementUpdateValue bson.D) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		objectID, err := primitive.ObjectIDFromHex(docId)
		if err != nil {
			logger.Log().Error("Invalid Record ID:", err)
			return false, err
		}
		selector := bson.D{
			{Key: "_id", Value: objectID},
			{Key: arrFieldPath + "." + elementSelectorName, Value: elementSelectorValue},
		}
		meta := repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(bson.D{}, &rc)
		update := bson.M{
			"$set": append(mongoUtility.FlattenBsonD(elementUpdateValue, arrFieldPath+".$."), meta...),
			"$inc": bson.M{"meta._version": 1},
		}
		// TODO apply security.
		result, err := repo.Collection.UpdateOne(ctx, selector, update)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with provided selector")
				return false, errors.New("record not found with provided selector")
			}
			logger.Log().Error("Failed to append element to array field:n", err)
			return false, err
		}
		if result.ModifiedCount == 0 {
			// Now again perform addToSet. in case entry already not exists.
			selectorForAts := bson.D{{Key: "_id", Value: objectID}}
			updateForAts := bson.M{"$addToSet": bson.D{{Key: arrFieldPath, Value: elementUpdateValue}}}
			_, err = repo.Collection.UpdateOne(ctx, selectorForAts, updateForAts)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					logger.Log().Error("Record not found with provided selector")
					return false, errors.New("record not found with provided selector")
				}
				logger.Log().Error("Failed to append element to array field:", err)
				return false, err
			}
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) RemoveFromArrayField(
	ctx context.Context,
	docId string,
	arrFieldPath string,
	elementSelectorName string,
	elementSelectorValue string) (bool, error) {
	return handleWrite(ctx, func(rc service.RequestContext) (bool, error) {
		objectID, err := primitive.ObjectIDFromHex(docId)
		if err != nil {
			logger.Log().Error("Invalid Record ID:", err)
			return false, err
		}
		selector := bson.D{
			{Key: "_id", Value: objectID},
		}
		meta := repo.RepositoryContext.MetaDataWriter.WriteUpdateMeta(bson.D{}, &rc)
		update := bson.M{
			"$set":  meta,
			"$pull": bson.M{arrFieldPath: bson.M{elementSelectorName: elementSelectorValue}},
			"$inc":  bson.M{"meta._version": 1},
		}
		// TODO apply security.
		_, err = repo.Collection.UpdateOne(ctx, selector, update)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with provided selector")
				return false, errors.New("record not found with provided selector")
			}
			logger.Log().Error("Failed to remove element from array field:", err)
			return false, err
		}
		return true, nil
	})
}

func (repo *MongoRepository[Entity]) FindById(ctx context.Context, id string) *Entity {
	return handleRead(ctx, func(rc service.RequestContext) *Entity {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Log().Error("Invalid ID: ", err)
			return nil
		}
		var document Entity
		//repo.RepositoryContext.DataSecurity.
		err = repo.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&document)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log().Error("Record not found with ID: ", id)
				return nil
			}
			logger.Log().Error("Failed to get document: ", err)
			return nil
		}
		return &document
	}, nil)
}

func (repo *MongoRepository[Entity]) Find(ctx context.Context, filters bson.D, offset int64, limit int64, sort bson.D) []*Entity {
	return handleRead(ctx, func(rc service.RequestContext) []*Entity {
		results := []*Entity{}
		if sort == nil {
			sort = bson.D{{Key: "meta._created", Value: -1}}
		}
		if limit <= 0 {
			limit = 25
		}
		options := &options.FindOptions{
			Skip:  &offset,
			Limit: &limit,
			Sort:  sort,
		}
		// TODO re-generate filter using seucrity check.
		cur, err := repo.Collection.Find(ctx, filters, options)
		if err != nil {
			logger.Log().Error("Error listing documents:", err)
			return results
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var document Entity
			err := cur.Decode(&document)
			if err != nil {
				logger.Log().Error("Error decoding document:", err)
			}
			results = append(results, &document)
		}
		if err := cur.Err(); err != nil {
			logger.Log().Error("Error iterating over documents:", err)
			return nil
		}
		return results
	}, []*Entity{})
}

func (repo *MongoRepository[Entity]) Count(ctx context.Context, filters bson.D) int64 {
	return handleRead(ctx, func(rc service.RequestContext) int64 {
		count, err := repo.Collection.CountDocuments(ctx, filters)
		if err != nil {
			logger.Log().Error("Error listing documents:", err)
			return 0
		}
		return count
	}, 0)
}

func (repo *MongoRepository[Entity]) Aggregate(ctx context.Context, pipelines []bson.D) []*bson.D {
	return handleRead(ctx, func(rc service.RequestContext) []*bson.D {
		results := []*bson.D{}
		// TODO re-generate filter using seucrity check.
		cur, err := repo.Collection.Aggregate(ctx, pipelines)
		if err != nil {
			logger.Log().Error("Error listing documents:", err)
			return results
		}
		defer cur.Close(ctx)
		if err := cur.All(ctx, &results); err != nil {
			logger.Log().Error("Error decoding document:", err)
			return nil
		}
		return results
	}, []*bson.D{})
}

func (repo *MongoRepository[Entity]) FindByIdPublic(ctx context.Context, id string) *Entity {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Log().Error("Invalid ID: %v\n", err)
		return nil
	}
	var document Entity
	//repo.RepositoryContext.DataSecurity.
	err = repo.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log().Error("Record not found with ID:", id)
			return nil
		}
		logger.Log().Error("Failed to get document:", err)
		return nil
	}
	return &document
}

func (repo *MongoRepository[Entity]) FindPublic(ctx context.Context, filters bson.D, offset int64, limit int64, sort bson.D) []*Entity {
	results := []*Entity{}
	if sort == nil {
		sort = bson.D{{Key: "meta._created", Value: -1}}
	}
	if limit <= 0 {
		limit = 25
	}
	options := &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
		Sort:  sort,
	}
	// TODO re-generate filter using seucrity check.
	cur, err := repo.Collection.Find(ctx, filters, options)
	if err != nil {
		logger.Log().Error("Error listing documents:", err)
		return results
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var document Entity
		err := cur.Decode(&document)
		if err != nil {
			logger.Log().Error("Error decoding document:", err)
		}
		results = append(results, &document)
	}
	if err := cur.Err(); err != nil {
		logger.Log().Error("Error iterating over documents:", err)
		return nil
	}
	return results
}

func (repo *MongoRepository[Entity]) CountPublic(ctx context.Context, filters bson.D) int64 {
	count, err := repo.Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		logger.Log().Error("Error listing documents:", err)
		return 0
	}
	return count
}
