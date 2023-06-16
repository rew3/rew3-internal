package utils

import (
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 * Convert given raw json message to BSON document.
 */
func JsonToBson(rawMessage json.RawMessage) (bson.M, error) {
	var bsonM bson.M
	err := bson.Unmarshal(rawMessage, &bsonM)
	if err != nil {
		return nil, err
	}
	return bsonM, nil
}

/*
 * Convert given BSON document to raw json message.
 */
func BsonToJson(bsonData bson.M) (json.RawMessage, error) {
	extendedJSON, err := bson.MarshalExtJSON(bsonData, true, false)
	if err != nil {
		return nil, err
	}
	var rawMessage json.RawMessage = extendedJSON
	return rawMessage, nil
}

/*
 * Convert given entity to Mongo BSON document.
 */
func EntityToBson[Entity any](entity *Entity, convertToObjectId bool) (bson.M, error) {
	bsonData, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}
	var doc bson.M
	err = bson.Unmarshal(bsonData, &doc)
	if err != nil {
		return nil, err
	}
	if convertToObjectId {
		if err := ConvertHexToObjectID(doc); err != nil {
			return nil, err
		}
	}
	return doc, nil
}

/*
 * Convert given mongo bson document to Entity.
 */
func BsonToEntity[Entity any](document bson.M, convertIdToHex bool) (*Entity, error) {
	if convertIdToHex {
		if err := ConvertHexToObjectID(document); err != nil {
			return nil, err
		}
	}
	var entity Entity
	bsonData, err := bson.Marshal(document)
	if err != nil {
		return nil, err
	}
	if err := bson.Unmarshal(bsonData, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

// Convert _id from string to primitive.ObjectID
func ConvertHexToObjectID(doc bson.M) error {
	if objectID, err := primitive.ObjectIDFromHex(doc["_id"].(string)); err == nil {
		doc["_id"] = objectID // Assign the ObjectID to the _id field
		return nil
	} else {
		return err
	}
}

// Convert _id from primitive.ObjectID to string.
func ConvertObjectIDToHex(doc bson.M) error {
	if data, ok := doc["_id"].(primitive.ObjectID); ok {
		doc["_id"] = data.Hex()
		return nil
	}
	return errors.New("_id field is not ObjectId")
}
