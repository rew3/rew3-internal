package json

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/rew3/rew3-internal/pkg/utils/logger"
)

/*
 * Convert given json data to given type struct.
 * Note: this will automatically convert RFC3339 format string date to time.Time. Any invalid string
 * date value will not be able to parse and throw error.
 *
 * Param:
 * jsonBytes - json data in bytes.
 */
func JsonToType[T any](jsonBytes []byte) (T, error) {
	var result T
	var inputMap map[string]interface{}
	err := json.Unmarshal(jsonBytes, &inputMap)
	if err != nil {
		logger.Log().Errorln("Error decoding JSON:", err)
		return result, err
	}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  result,
		Squash:  true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339), // Convert string to time.Time using RFC3339 format
		),
	})
	if err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return result, err
	}
	if err := decoder.Decode(inputMap); err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return result, err
	}
	return result, nil
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
		logger.Log().Errorln("Error decoding JSON:", err)
		return inputMap, err
	}
	return inputMap, nil
}

/*
 * Convert given json data to given type struct.
 * Note: this will automatically convert RFC3339 format string date to time.Time. Any invalid string
 * date value will not be able to parse and throw error.
 *
 * Param:
 * data - data to convert to json.
 */
func TypeToJson[T any](data T) ([]byte, error) {
	outputMap := make(map[string]interface{})
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  outputMap,
		Squash:  true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339), // Convert string to time.Time using RFC3339 format
		),
	})
	if err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return nil, err
	}
	// Convert map to JSON
	jsonData, err := json.Marshal(outputMap)
	if err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return nil, err
	}
	return jsonData, nil
}

/*
 * Convert given json data to given type struct.
 * Note: this will automatically convert RFC3339 format string date to time.Time. Any invalid string
 * date value will not be able to parse and throw error.
 *
 * Param:
 * data - data to convert to map.
 */
func TypeToMap[T any](data T) (map[string]interface{}, error) {
	outputMap := make(map[string]interface{})
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  outputMap,
		Squash:  true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339), // Convert string to time.Time using RFC3339 format
		),
	})
	if err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		logger.Log().Errorln("Error decoding data:", err)
		return nil, err
	}
	return outputMap, nil
}
