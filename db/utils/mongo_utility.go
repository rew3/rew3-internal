package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/rew3/rew3-internal/pkg/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 * Convert given raw json message to bson.M document.
 */
func JsonToBsonM(rawMessage json.RawMessage) (bson.M, error) {
	bsonM := bson.M{}
	err := bson.Unmarshal(rawMessage, &bsonM)
	if err != nil {
		return nil, err
	}
	return bsonM, nil
}

/*
 * Convert given raw json message to bson.D document.
 */
func JsonToBsonD(rawMessage json.RawMessage) (bson.D, error) {
	bsonD := bson.D{}
	err := bson.Unmarshal(rawMessage, &bsonD)
	if err != nil {
		return nil, err
	}
	return bsonD, nil
}

/*
 * Convert given bson.M document to raw json message.
 */
func BsonMToJson(bsonData bson.M) (json.RawMessage, error) {
	extendedJSON, err := bson.MarshalExtJSON(bsonData, true, false)
	if err != nil {
		return nil, err
	}
	var rawMessage json.RawMessage = extendedJSON
	return rawMessage, nil
}

/*
 * Convert given bson.D document to raw json message.
 */
func BsonDToJson(bsonData bson.D) (json.RawMessage, error) {
	extendedJSON, err := bson.MarshalExtJSON(bsonData, true, false)
	if err != nil {
		return nil, err
	}
	var rawMessage json.RawMessage = extendedJSON
	return rawMessage, nil
}

/*
 * Convert given entity to Mongo bson.M document.
 * If noCache true proivided - default decoder is used else bson.Unmarshal used where mongo will apply internal cache.
 */
func EntityToBsonM[Entity any](entity *Entity, convertToObjectId bool, removeEmptyFields bool, noCache bool) (bson.M, error) {
	bsonData, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}
	doc := bson.M{}
	if noCache {
		dec, err := bson.NewDecoder(bsonrw.NewBSONDocumentReader(bsonData))
		if err != nil {
			return nil, err
		}
		dec.SetRegistry(bson.DefaultRegistry)
		if err = dec.Decode(&doc); err != nil {
			return nil, err
		}
	} else {
		err := bson.Unmarshal(bsonData, &doc)
		if err != nil {
			return nil, err
		}
	}
	if convertToObjectId {
		if err := BsonMHexToObjectID(doc); err != nil {
			return nil, err
		}
	}
	if removeEmptyFields {
		doc = RemoveEmptyBsonMFields(doc)
	}
	return doc, nil
}

/*
 * Convert given entity to Mongo bson.D document.
 * If noCache true proivided - default decoder is used else bson.Unmarshal used where mongo will apply internal cache.
 */
func EntityToBsonD[Entity any](entity *Entity, convertToObjectId bool, removeEmptyFields bool, noCache bool) (bson.D, error) {
	bsonData, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}
	doc := bson.D{}
	if noCache {
		dec, err := bson.NewDecoder(bsonrw.NewBSONDocumentReader(bsonData))
		if err != nil {
			return nil, err
		}
		dec.SetRegistry(bson.DefaultRegistry)
		if err = dec.Decode(&doc); err != nil {
			return nil, err
		}
	} else {
		err := bson.Unmarshal(bsonData, &doc)
		if err != nil {
			return nil, err
		}
	}
	if convertToObjectId {
		if err := BsonDHexToObjectID(doc); err != nil {
			return nil, err
		}
	}
	if removeEmptyFields {
		doc = RemoveEmptyBsonDFields(doc)
	}
	return doc, nil
}

/*
 * Convert given mongo bson document to Entity.
 */
func BsonMToEntity[Entity any](document bson.M, convertIdToHex bool, val *Entity, noCache bool) error {
	if convertIdToHex {
		if err := BsonMObjectIDToHex(document); err != nil {
			return err
		}
	}
	bsonData, err := bson.Marshal(document)
	if err != nil {
		return err
	}
	if noCache {
		dec, err := bson.NewDecoder(bsonrw.NewBSONDocumentReader(bsonData))
		if err != nil {
			return err
		}
		dec.SetRegistry(bson.DefaultRegistry)
		if err = dec.Decode(&val); err != nil {
			return err
		}
	} else {
		err := bson.Unmarshal(bsonData, &val)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
 * Convert given mongo bson document to Entity.
 */
func BsonDToEntity[Entity any](document bson.D, convertIdToHex bool, val *Entity, noCache bool) error {
	if convertIdToHex {
		if err := BsonDObjectIDToHex(document); err != nil {
			return err
		}
	}
	bsonData, err := bson.Marshal(document)
	if err != nil {
		return err
	}
	if noCache {
		dec, err := bson.NewDecoder(bsonrw.NewBSONDocumentReader(bsonData))
		if err != nil {
			return err
		}
		dec.SetRegistry(bson.DefaultRegistry)
		if err = dec.Decode(&val); err != nil {
			return err
		}
	} else {
		err := bson.Unmarshal(bsonData, &val)
		if err != nil {
			return err
		}
	}
	return nil
}

// Convert _id from string to primitive.ObjectID
func BsonMHexToObjectID(doc bson.M) error {
	if objectID, err := primitive.ObjectIDFromHex(doc["_id"].(string)); err == nil {
		doc["_id"] = objectID // Assign the ObjectID to the _id field
		return nil
	} else {
		return err
	}
}
func BsonDHexToObjectID(doc bson.D) error {
	for i, elem := range doc {
		if elem.Key == "_id" {
			if hexID, ok := elem.Value.(string); ok {
				objectID, err := primitive.ObjectIDFromHex(hexID)
				if err != nil {
					return err
				}
				doc[i].Value = objectID
				return nil
			}
		}
	}
	return errors.New("_id field not found or is not a valid hexadecimal string")
}

/*
 * Convert _id from primitive.ObjectID to string of given bson.M document.
 */
func BsonMObjectIDToHex(doc bson.M) error {
	if data, ok := doc["_id"].(primitive.ObjectID); ok {
		doc["_id"] = data.Hex()
		return nil
	}
	return errors.New("_id field is not ObjectId")
}

/*
 * Convert _id from primitive.ObjectID to string of given bson.D document.
 */
func BsonDObjectIDToHex(doc bson.D) error {
	for i := range doc {
		if doc[i].Key == "_id" {
			if objectID, ok := doc[i].Value.(primitive.ObjectID); ok {
				doc[i].Value = objectID.Hex()
				return nil
			}
		}
	}
	return errors.New("_id field is not found or is not an ObjectId")
}

/*
 * Remove empty fields of bson.M document of type document.
 */
func RemoveEmptyBsonMFields(bsonDoc bson.M) bson.M {
	filteredBSON := bson.M{}
	for key, value := range bsonDoc {
		if subDoc, ok := value.(bson.M); ok {
			subDoc = RemoveEmptyBsonMFields(subDoc) // Recursively check nested fields
			if len(subDoc) == 0 {
				continue // Skip empty nested fields
			}
			filteredBSON[key] = subDoc
		} else {
			filteredBSON[key] = value
		}
	}
	return filteredBSON
}

/*
 * Remove empty fields of bson.D document of type document.
 */
func RemoveEmptyBsonDFields(bsonDoc bson.D) bson.D {
	filteredBSON := bson.D{}
	for _, elem := range bsonDoc {
		if subDoc, ok := elem.Value.(bson.D); ok {
			subDoc = RemoveEmptyBsonDFields(subDoc) // Recursively check nested fields
			if len(subDoc) == 0 {
				continue // Skip empty nested fields
			}
			elem.Value = subDoc
		}
		filteredBSON = append(filteredBSON, elem)
	}
	return filteredBSON
}

/*
 * Get the value of bson.M document field.
 * Note: In case of error of value not found, default value is returned.
 */
func GetBsonMFieldValueWithDefault[T any](doc bson.M, field string, defaultValue T) T {
	if value, err := GetBsonMFieldValue[T](doc, field); err == nil {
		return value
	} else {
		return defaultValue
	}
}

/*
 * Get the value of bson.D document field.
 * Note: In case of error of value not found, default value is returned.
 */
func GetBsonDFieldValueWithDefault[T any](doc bson.D, field string, defaultValue T) T {
	if value, err := GetBsonDFieldValue[T](doc, field); err == nil {
		return value
	} else {
		return defaultValue
	}
}

/*
 * Get the value of bson.M document field.
 * Note: You can provide field name in dot notation format.
 */
func GetBsonMFieldValue[T any](doc bson.M, field string) (T, error) {
	keys := strings.Split(field, ".")
	length := len(keys)
	if length == 0 {
		var value T
		return value, errors.New("invalid field key provided")
	} else if length == 1 {
		if data, ok := doc[keys[0]].(T); ok {
			return data, nil
		} else {
			var value T
			return value, errors.New("the field value if did not match given type")
		}
	} else {
		key := keys[0]
		if data, ok := doc[key].(bson.M); ok {
			return GetBsonMFieldValue[T](data, strings.Join(keys[1:], "."))
		} else {
			var value T
			return value, fmt.Errorf("field `%s` is not nested type", key)
		}
	}
}

/*
 * Get the value of bson.D document field.
 * Note: You can provide field name in dot notation format.
 */
func GetBsonDFieldValue[T any](doc bson.D, field string) (T, error) {
	keys := strings.Split(field, ".")
	length := len(keys)
	if length == 0 {
		var value T
		return value, errors.New("invalid field key provided")
	} else if length == 1 {
		for _, elem := range doc {
			if elem.Key == keys[0] {
				if data, ok := elem.Value.(T); ok {
					return data, nil
				} else {
					var value T
					return value, errors.New("the field value did not match the given type")
				}
			}
		}
		var value T
		return value, fmt.Errorf("field `%s` not found", keys[0])
	} else {
		key := keys[0]
		for _, elem := range doc {
			if elem.Key == key {
				if data, ok := elem.Value.(bson.D); ok {
					return GetBsonDFieldValue[T](data, strings.Join(keys[1:], "."))
				} else {
					var value T
					return value, fmt.Errorf("field `%s` is not a nested type", key)
				}
			}
		}
		var value T
		return value, fmt.Errorf("field `%s` not found", key)
	}
}

/*
 * Set given value to the provided bson.D document field.
 * Note: You can provide field name in dot notation format.
 */
func SetBsonDFieldValue(doc bson.D, field string, value interface{}) error {
	keys := strings.Split(field, ".")
	length := len(keys)
	if length == 0 {
		return errors.New("invalid field key provided")
	} else if length == 1 {
		for i := range doc {
			if doc[i].Key == keys[0] {
				doc[i].Value = value
				return nil
			}
		}
		return fmt.Errorf("field `%s` not found", keys[0])
	} else {
		key := keys[0]
		for i := range doc {
			if doc[i].Key == key {
				if data, ok := doc[i].Value.(bson.D); ok {
					return SetBsonDFieldValue(data, strings.Join(keys[1:], "."), value)
				} else {
					return fmt.Errorf("field `%s` is not a nested type", key)
				}
			}
		}
		return fmt.Errorf("field `%s` not found", key)
	}
}

/*
 * Add New field value for given bson.D docuemnt.
 */
func AddBsonDNewFieldValue(doc bson.D, field string, value interface{}) bson.D {
	element := bson.E{Key: field, Value: value}
	return append(doc, element)
}

/*
 * Remove given field from bson.M document.
 * Note, nested level field are not supported i.e. no dot notation field name accepted.
 */
func RemoveFieldFromBsonM(doc bson.M, field string) {
	delete(doc, field)
}

/*
 * Remove given field from bson.D document.
 * Note, nested level field are not supported i.e. no dot notation field name accepted.
 */
func RemoveFieldFromBsonD(doc bson.D, field string) bson.D {
	filtered := bson.D{}
	for _, elem := range doc {
		if elem.Key != field {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

/*
 * Write new object ID field to given bson.D document.
 */
func WriteObjectID(doc bson.D) (primitive.ObjectID, bson.D) {
	id := primitive.NewObjectID()
	prepended := bson.D{bson.E{Key: "_id", Value: id}}
	prepended = append(prepended, doc...)
	return id, prepended
}

func PrintJson(document bson.D) {
	// Convert bson.D to JSON bytes
	jsonBytes, err := bson.MarshalExtJSON(document, true, false)
	if err != nil {
		logger.Log().Error("Error marshaling BSON to JSON: ", err.Error())
		return
	}
	// Convert JSON bytes to string
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}

// FlattenBSON recursively flattens a BSON document with nested fields using dot notation.
func FlattenBsonD(doc bson.D, prefixToAdd string) bson.D {
	flattened := bson.D{}
	flattenHelper(&flattened, "", doc, prefixToAdd)
	return flattened
}

func flattenHelper(flattened *bson.D, prefix string, doc bson.D, prefixToAdd string) {
	tmpFlattened := *flattened
	for _, elem := range doc {
		key := elem.Key
		if prefix != "" {
			key = fmt.Sprintf("%s.%s", prefix, key)
		}
		if prefixToAdd != "" {
			key = fmt.Sprintf("%s%s", prefixToAdd, key)
		}
		if reflect.TypeOf(elem.Value) == reflect.TypeOf(bson.D{}) {
			subdoc, ok := elem.Value.(bson.D)
			if ok {
				flattenHelper(&tmpFlattened, key, subdoc, "")
			}
		} else {
			tmpFlattened = append(tmpFlattened, bson.E{Key: key, Value: elem.Value})
		}
	}
	*flattened = tmpFlattened
}

/**
 * Resolve raw string json and return bson.D document.
 * Note: if invalid json, empty bson document is returned.
 */
func RawJsonStringToBson(rawValue string) bson.D {
	if rawValue == "" {
		return bson.D{}
	}
	bsonDoc := bson.D{}
	err := bson.UnmarshalExtJSON([]byte(rawValue), true, &bsonDoc)
	if err != nil {
		logger.Log().Error("Unable to resolve sort value: ", err.Error())
		return bson.D{}
	} else {
		return bsonDoc
	}
}
