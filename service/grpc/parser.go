package grpc

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/tidwall/gjson"
)

/**
 * Parser to parse payload data (request and response)
 */
type PayloadWrapper struct {
	ParsedJson gjson.Result
	Original   json.RawMessage
}

func WrapPayload(data json.RawMessage) *PayloadWrapper {
	parsed := gjson.ParseBytes(data)
	return &PayloadWrapper{ParsedJson: parsed, Original: data}
}

/**
 * Parse value at given path to provided type.
 */
func ParsePayload[T any](payload *PayloadWrapper, path string) (T, error) {
	var data T
	rawJson := payload.ParsedJson.Get(path).Raw
	if rawJson == "" {
		return data, errors.New("[ParsePayload] parsing error: invalid data provided")
	}
	err := json.Unmarshal([]byte(rawJson), &data)
	if err != nil {
		log.Println("[ParsePayload] Parsing Error:", err.Error())
		return data, errors.New("Parsing Error: " + err.Error())
	}
	return data, nil
}

/**
 * Parse given type to json.
 */
func ToJson[T any](data T) (json.RawMessage, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("[ToJson] Parsing Error:", err.Error())
		return nil, err
	}
	return json.RawMessage(jsonData), nil
}

/**
 * Parse given raw json to type.
 */
func ToType[T any](data json.RawMessage) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Println("[ToType] Parsing Error:", err.Error())
		return result, err
	}
	return result, nil
}
