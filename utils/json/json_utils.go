package json

import (
	"encoding/json"

	"github.com/rew3/rew3-pkg/utils/logger"
)

/*
 * Json Serialization Utility.
 * Note: this will automatically convert time.Time to RFC3339 format string date.
 */

/*
 * Convert given json data to given type struct.
 *
 * Param:
 * jsonBytes - json data in bytes.
 */
func JsonToType[T any](jsonBytes []byte) (T, error) {
	var result T
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		logger.Log().Errorln("Error decoding JSON:", err)
		return result, err
	}
	return result, nil
}

/*
 * Convert given json data to given type struct.
 *
 * Param:
 * data - data to convert to json.
 */
func TypeToJson[T any](data T) ([]byte, error) {
	// Convert struct to JSON-encoded byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log().Errorln("Error encoding data to Json:", err)
		return nil, err
	}
	return jsonData, nil
}

/*
 * Convert given json data to map.
 *
 * Param:
 * jsonBytes - json data in bytes.
 */
func JsonToMap(jsonBytes []byte) (map[string]interface{}, error) {
	var inputMap map[string]interface{}
	err := json.Unmarshal(jsonBytes, &inputMap)
	if err != nil {
		logger.Log().Errorln("Error decoding data to Map:", err)
		return inputMap, err
	}
	return inputMap, nil
}

/*
 * Convert given json data to map.
 *
 * Param:
 * jsonBytes - json data in bytes.
 */
func MapToJson(input map[string]interface{}) ([]byte, error) {
	// Convert struct to JSON-encoded byte slice
	jsonData, err := json.Marshal(input)
	if err != nil {
		logger.Log().Errorln("Error encoding data to Json:", err)
		return nil, err
	}
	return jsonData, nil
}

/*
 * Convert given struct data to map.
 *
 * Param:
 * data - data to convert to map.
 */
func TypeToMap[T any](data T) (map[string]interface{}, error) {
	// Convert struct to JSON-encoded byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log().Errorln("Error encoding data to Map:", err)
		return nil, err
	}
	// Unmarshal JSON data into a map
	var resultMap map[string]interface{}
	err = json.Unmarshal(jsonData, &resultMap)
	if err != nil {
		logger.Log().Errorln("Error encoding data to Map:", err)
		return nil, err
	}
	return resultMap, nil
}

/*
 * Convert given map to struct.
 *
 * Param:
 * data - data to convert to map.
 */
func MapToType[T any](data map[string]interface{}) (T, error) {
	var result T
	// Convert map to JSON-encoded byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log().Errorln("Error decoding data to struct:", err)
		return result, err
	}

	// Unmarshal JSON data into the custom type
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		logger.Log().Errorln("Error decoding data to struct:", err)
		return result, err
	}
	return result, nil
}
