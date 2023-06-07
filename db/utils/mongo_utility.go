package utils

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
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
func EntityToBson[Entity any](entity *Entity) (bson.M, error) {
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
