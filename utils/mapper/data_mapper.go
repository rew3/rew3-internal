package mapper

import (
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/rew3/rew3-pkg/utils/logger"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

/*
 * Convert Model type to Proto Type.
 * Note: Proto Message type P must be pointer type i.e. *ProtoMessage e.g. *ContactProto.
 */
func MapModelToPB[M any, P proto.Message](model M, defaultValue P) (P, error) {
	// Converting model to Map.
	maps, err := buildMapFromModel(model)
	if err != nil {
		logger.Log().Error("MapModelToProto: Error while converting from Modal to Maps:", err)
		return defaultValue, err
	}
	// Converting maps to proto.
	proto, err := buildProtoFromMap(maps, defaultValue)
	if err != nil {
		logger.Log().Error("MapModelToProto: Error while converting from Maps to Proto:", err)
		return defaultValue, err
	}
	return proto, nil
}

/*
 * Convert Proto type to Model Type.
 * If trimString provided true, all string field with leading/trailing whitespaces will be trimed.
 * Note: Proto Message type P must be pointer type i.e. *ProtoMessage e.g. *ContactProto.
 */
func MapModelFromPB[M any, P proto.Message](proto P, defaultValue M, trimString bool) (M, error) {
	// Converting proto to Map.
	maps, err := buildMapFromProto(proto)
	if err != nil {
		logger.Log().Error("MapModelFromProto: Error while converting from Proto to Maps:", err)
		return defaultValue, err
	}
	if trimString {
		recursivelyTrimStrings(maps)
	}
	// Converting maps to Model.
	result, err := buildModelFromMap(maps, &defaultValue)
	if err != nil {
		logger.Log().Error("MapModelFromProto: Error while converting from Maps to Model:", err)
		return defaultValue, err
	}
	return *result, nil
}

// recursivelyTrimStrings trims string values in a nested map structure
func recursivelyTrimStrings(data interface{}) {
	switch d := data.(type) {
	case map[string]interface{}:
		for key, value := range d {
			if strValue, ok := value.(string); ok {
				trimmedValue := strings.TrimSpace(strValue)
				d[key] = trimmedValue
			} else {
				recursivelyTrimStrings(value)
			}
		}
	case []interface{}:
		for i, v := range d {
			recursivelyTrimStrings(v)
			d[i] = v
		}
	}
}

/*
 * INTERNAL: Convert Proto type to Map.
 */
func buildMapFromProto(pbMessage proto.Message) (map[string]interface{}, error) {
	// Create a new JSON marshaller
	marshaler := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: false,
	}
	// Marshal the Protobuf message to JSON
	jsonBytes, err := marshaler.Marshal(pbMessage)
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON to a map
	var resultMap map[string]interface{}
	strings.TrimSpace(string(jsonBytes))
	if err = json.Unmarshal(jsonBytes, &resultMap); err != nil {
		return nil, err
	}
	return resultMap, nil
}

/*
 * INTERNAL: Convert Map to Proto type.
 *
 * Param:
 * inputMap - to convert values in map.
 * pbStruct - Values will be read to this proto struct.
 */
func buildProtoFromMap[T proto.Message](inputMap map[string]interface{}, pbStruct T) (T, error) {
	// Marshal the input map to JSON
	jsonData, err := json.Marshal(inputMap)
	if err != nil {
		return pbStruct, err
	}
	// Unmarshal the JSON data into the Protobuf struct
	if err := protojson.Unmarshal(jsonData, pbStruct); err != nil {
		return pbStruct, err
	}
	return pbStruct, nil
}

/*
 * INTERNAL: Convert Model type to Map.
 */
func buildMapFromModel(model interface{}) (map[string]interface{}, error) {
	var resultMap map[string]interface{}
	jsonData, err := json.Marshal(model)
	if err != nil {
		return resultMap, err
	}
	if err := json.Unmarshal(jsonData, &resultMap); err != nil {
		return resultMap, err
	}
	return resultMap, nil
}

/*
 * INTERNAL: Convert Map to Model type.
 * Note: this will automatically convert RFC3339 format string date to time.Time. Any invalid string
 * date value will not be able to parse and throw error.
 *
 * Param:
 * inputMap - to convert values in map.
 * model - pointer to the struct that will contain the decoded value.
 */
func buildModelFromMap[T any](inputMap map[string]interface{}, model T) (T, error) {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  model,
		Squash:  true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339), // Convert string to time.Time using RFC3339 format
		),
	})
	if err != nil {
		return model, err
	}
	if err := decoder.Decode(inputMap); err != nil {
		return model, err
	}
	return model, nil
}
