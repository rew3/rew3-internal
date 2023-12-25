package message

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
)

/**
 * Coded to serialize and deserialize message.
 * ST - to serialize type, DT - to deserialize type.
 */
type Codec[ST any, DT any] struct {
	Serializer   Serializer[ST, DT]
	Deserializer Deserializer[ST, DT]
}

/**
 * Create default codec.
 * This coded will serialize/deserialize given input into/from bytes.
 */
func DefaultCodec[T any]() *Codec[T, []byte] {
	return &Codec[T, []byte]{
		Serializer:   MQSerializer[T]{},
		Deserializer: MQDeserializer[T]{},
	}
}

/**
 * Interface for Serializer.
 */
type Serializer[I any, O any] interface {
	Serialize(input I) (O, error)
}

/**
 * Interface for De-serializer.
 */
type Deserializer[I any, O any] interface {
	Deserialize(input O) (I, error)
}

/**
 * Default Serializer.
 * Serialize given input type data to string UTF-8.
 */
type MQSerializer[I any] struct {
	Serializer[I, []byte]
}

/**
 * Serialize given input into byte streams.
 */
func (s MQSerializer[I]) Serialize(input I) ([]byte, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		logger.Log().Println("ERROR serialization: ", err)
		return nil, err
	}
	return jsonBytes, nil
}

/**
 * Default De-Serializer.
 * Deserialize string input raw data into given input type I.
 */
type MQDeserializer[O any] struct {
	Deserializer[O, []byte]
}

/**
 * De-Serialize to output type.
 */
func (s MQDeserializer[O]) Deserialize(input []byte) (O, error) {
	var result O
	err := json.Unmarshal(input, &result)
	if err != nil {
		logger.Log().Println("ERROR de-serialization: ", err)
		return result, err
	}
	return result, nil
}

/*
 * Convert given input into map.
 */
func ToMap(input interface{}) (map[string]interface{}, error) {
	var resultMap map[string]interface{}
	jsonData, err := json.Marshal(input)
	if err != nil {
		return resultMap, err
	}
	if err := json.Unmarshal(jsonData, &resultMap); err != nil {
		return resultMap, err
	}
	return resultMap, nil
}

/*
 * Convert given map into output type T.
 * Note: this will automatically convert RFC3339 format string date to time.Time. Any invalid string
 * date value will not be able to parse and throw error.
 */
func FromMap[T any](inputMap map[string]interface{}) (T, error) {
	var output T
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  output,
		Squash:  true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339), // Convert string to time.Time using RFC3339 format
		),
	})
	if err != nil {
		return output, err
	}
	if err := decoder.Decode(inputMap); err != nil {
		return output, err
	}
	return output, nil
}
