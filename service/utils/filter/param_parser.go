package filter

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/rew3/rew3-internal/service/shared/request"
)

/**
 * Filter parameter Parser utility.
 */

type FilterOperator string
type FilterValueType string

const (
	NONE FilterOperator = ""

	// For All Data Types.
	IS        FilterOperator = "-is"
	IS_NOT    FilterOperator = "-is-not"
	EMPTY     FilterOperator = "-empty"
	NOT_EMPTY FilterOperator = "-not-empty"
	IN        FilterOperator = "-in"
	NOT_IN    FilterOperator = "-not-in"

	// For String.
	CONTAINS        FilterOperator = "-contains"
	NOT_CONTAINS    FilterOperator = "-not-contains"
	STARTS_WITH     FilterOperator = "-starts-with"
	NOT_STARTS_WITH FilterOperator = "-not-starts-with"
	ENDS_WITH       FilterOperator = "-ends-with"
	NOT_ENDS_WITH   FilterOperator = "-not-ends-with"

	// For Numbers and Date (applicable to String too)
	EQ                     FilterOperator = "-eq"
	NOT_EQ                 FilterOperator = "-not-eq"
	LESS_THAN              FilterOperator = "-lt"
	NOT_LESS_THAN          FilterOperator = "-not-lt"
	LESS_THAN_EQUAL        FilterOperator = "-lt-eq"
	NOT_LESS_THAN_EQUAL    FilterOperator = "-not-lt-eq"
	GREATER_THAN           FilterOperator = "-gt"
	NOT_GREATER_THAN       FilterOperator = "-not-gt"
	GREATER_THAN_EQUAL     FilterOperator = "-gt-eq"
	NOT_GREATER_THAN_EQUAL FilterOperator = "-not-gt-eq"

	// For Number, Date.
	RANGE        FilterOperator = "-range"
	NOT_IN_RANGE FilterOperator = "-not-in-range"

	// Date Type.
	INTEGER FilterValueType = "integer"
	LONG    FilterValueType = "long"
	FLOAT   FilterValueType = "float"
	DOUBLE  FilterValueType = "double"
	BOOLEAN FilterValueType = "boolean"
	STRING  FilterValueType = "string"
	ISODATE FilterValueType = "iso_date"
	UNKNOWN FilterValueType = "unknown"
)

type Filter struct {
	Field    string
	Value    interface{}
	Operator FilterOperator
}

/**
 * Parse filter parameter.
 */
func ParseFilterParams(param request.RequestParam) ([]Filter, error) {
	if param.Filters == "" {
		return []Filter{}, nil
	}
	filters := param.Filters
	splits := strings.Split(filters, "&")
	result := []Filter{}
	for _, f := range splits {
		filter, err := parseFilter(f)
		if err != nil {
			return nil, err
		}
		result = append(result, filter)
	}
	return result, nil
}

/**
 * Parse Filter for given raw filter string.
 */
func parseFilter(filter string) (Filter, error) {
	splits := strings.Split(filter, "=")
	if len(splits) == 2 {
		key := splits[0]
		value := splits[1]
		dataType, fieldName, op := parseFieldAndOperator(key)
		if op == RANGE || op == NOT_IN_RANGE {
			parsedValue, err := parseRangeValue(dataType, value)
			if err != nil {
				return Filter{}, fmt.Errorf("request parameter error. invalid range value `%s` is provided for parameter `%s`", value, key)
			}
			return Filter{
				Field:    fieldName,
				Value:    parsedValue,
				Operator: op,
			}, nil
		} else if op == IN || op == NOT_IN {
			parsedValue, err := parseListValue(dataType, value)
			if err != nil {
				return Filter{}, fmt.Errorf("request parameter error. invalid list value `%s` is provided for parameter `%s`", value, key)
			}
			return Filter{
				Field:    fieldName,
				Value:    parsedValue,
				Operator: op,
			}, nil
		} else {
			return Filter{
				Field:    fieldName,
				Value:    parseValue(dataType, value),
				Operator: op,
			}, nil
		}
	} else {
		key := splits[0]
		_, fieldName, op := parseFieldAndOperator(key)
		if op == EMPTY || op == NOT_EMPTY {
			return Filter{
				Field:    fieldName,
				Value:    nil,
				Operator: op,
			}, nil
		} else {
			return Filter{}, fmt.Errorf("request parameter error. invalid/empty value is provided for parameter `%s`", key)
		}
	}
}

/**
 * Parse Field and Operator.
 */
func parseFieldAndOperator(key string) (FilterValueType, string, FilterOperator) {
	dateType := UNKNOWN
	remaining := key
	parts := strings.Split(key, ":")
	if len(parts) == 2 {
		dateType = FilterValueType(parts[0])
		remaining = parts[1]
	}
	index := strings.Index(remaining, "-")
	if index < 0 || remaining[:index] == "" {
		// No filter operator type applied.
		return dateType, remaining, NONE
	}
	fieldName := remaining[:index]
	operator := remaining[index:]
	switch operator {
	case string(IS):
		return dateType, fieldName, IS
	case string(IS_NOT):
		return dateType, fieldName, IS_NOT
	case string(EMPTY):
		return dateType, fieldName, EMPTY
	case string(NOT_EMPTY):
		return dateType, fieldName, NOT_EMPTY
	case string(IN):
		return dateType, fieldName, IN
	case string(NOT_IN):
		return dateType, fieldName, NOT_IN
	case string(CONTAINS):
		return dateType, fieldName, CONTAINS
	case string(NOT_CONTAINS):
		return dateType, fieldName, NOT_CONTAINS
	case string(STARTS_WITH):
		return dateType, fieldName, STARTS_WITH
	case string(NOT_STARTS_WITH):
		return dateType, fieldName, NOT_STARTS_WITH
	case string(ENDS_WITH):
		return dateType, fieldName, ENDS_WITH
	case string(NOT_ENDS_WITH):
		return dateType, fieldName, NOT_ENDS_WITH
	case string(EQ):
		return dateType, fieldName, EQ
	case string(NOT_EQ):
		return dateType, fieldName, NOT_EQ
	case string(LESS_THAN):
		return dateType, fieldName, LESS_THAN
	case string(NOT_LESS_THAN):
		return dateType, fieldName, NOT_LESS_THAN
	case string(LESS_THAN_EQUAL):
		return dateType, fieldName, LESS_THAN_EQUAL
	case string(NOT_LESS_THAN_EQUAL):
		return dateType, fieldName, NOT_LESS_THAN_EQUAL
	case string(GREATER_THAN):
		return dateType, fieldName, GREATER_THAN
	case string(NOT_GREATER_THAN):
		return dateType, fieldName, NOT_GREATER_THAN
	case string(GREATER_THAN_EQUAL):
		return dateType, fieldName, GREATER_THAN_EQUAL
	case string(NOT_GREATER_THAN_EQUAL):
		return dateType, fieldName, NOT_GREATER_THAN_EQUAL
	case string(RANGE):
		return dateType, fieldName, RANGE
	case string(NOT_IN_RANGE):
		return dateType, fieldName, NOT_IN_RANGE
	default:
		return dateType, fieldName, NONE
	}
}

/**
 * Parse Value to appropriate data type.
 */
func parseValue(dataType FilterValueType, value string) interface{} {
	if dataType == UNKNOWN {
		if parsed, err := parseDateTimeValue(value); err == nil {
			return parsed
		} else {
			return parsePrimitiveValue(value)
		}
	} else {
		return parsePrimitiveValueToType(dataType, value)
	}
}

/**
 * Parse range value.
 */
func parseRangeValue(dataType FilterValueType, tupleStr string) ([2]interface{}, error) {
	tupleStr = strings.Trim(tupleStr, "[]")
	elements := strings.Split(tupleStr, ",")
	if len(elements) != 2 {
		return [2]interface{}{}, fmt.Errorf("invalid tuple format: expected 2 elements, got %d", len(elements))
	}
	var parsedElements [2]interface{}
	for i, element := range elements {
		element = strings.TrimSpace(element)
		if dataType == UNKNOWN {
			parsedElements[i] = parsePrimitiveValue(element)
		} else {
			parsedElements[i] = parsePrimitiveValueToType(dataType, element)
		}
	}
	return parsedElements, nil
}

/**
 * Parse list value.
 */
func parseListValue(dataType FilterValueType, listStr string) ([]interface{}, error) {
	listStr = strings.Trim(listStr, "[]")
	items := strings.Split(listStr, ",")
	var parsedItems []interface{}
	for _, item := range items {
		item = strings.TrimSpace(item)
		if dataType == UNKNOWN {
			parsedItems = append(parsedItems, parsePrimitiveValue(item))
		} else {
			parsedItems = append(parsedItems, parsePrimitiveValueToType(dataType, item))
		}
	}
	return parsedItems, nil
}

/**
 * Parse datetime value.
 */
func parseDateTimeValue(dateStr string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date string: %v", err)
	}
	return parsedTime, nil
}

/**
 * Parse primitive value with auto type inference.
 */
func parsePrimitiveValue(value string) interface{} {
	// Try parsing as bool
	if parsedValue, err := strconv.ParseBool(value); err == nil {
		return parsedValue
	}

	// Try parsing as int
	if parsedValue, err := strconv.Atoi(value); err == nil {
		return parsedValue
	}

	// Try parsing as int32
	if parsedValue, err := strconv.ParseInt(value, 10, 32); err == nil {
		return int32(parsedValue)
	}

	// Try parsing as int64
	if parsedValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return parsedValue
	}

	// Try parsing as float32
	if parsedValue, err := strconv.ParseFloat(value, 32); err == nil {
		return float32(parsedValue)
	}

	// Try parsing as float64
	if parsedValue, err := strconv.ParseFloat(value, 64); err == nil {
		return parsedValue
	}

	// Try parsing as big.Int
	bigIntValue := new(big.Int)
	if _, success := bigIntValue.SetString(value, 10); success {
		return bigIntValue
	}

	// Try parsing as time.Time
	if date, err := time.Parse(time.RFC3339, value); err == nil {
		return date
	}

	// Assume it as string if no match is found
	return value
}

/**
 * Parse primitive value to given data type.
 */
func parsePrimitiveValueToType(dataType FilterValueType, value string) interface{} {
	switch dataType {
	case INTEGER:
		// Try parsing as int32
		if parsedValue, err := strconv.ParseInt(value, 10, 32); err == nil {
			return int32(parsedValue)
		}
	case LONG:
		// Try parsing as int64
		if parsedValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return parsedValue
		}
	case FLOAT:
		// Try parsing as float32
		if parsedValue, err := strconv.ParseFloat(value, 32); err == nil {
			return float32(parsedValue)
		}
	case DOUBLE:
		// Try parsing as float64
		if parsedValue, err := strconv.ParseFloat(value, 64); err == nil {
			return parsedValue
		}
	case BOOLEAN:
		// Try parsing as bool
		if parsedValue, err := strconv.ParseBool(value); err == nil {
			return parsedValue
		}
	case ISODATE:
		// Try parsing as time.Time
		if date, err := time.Parse(time.RFC3339, value); err == nil {
			return date
		}
	default:
		return value
	}
	return value
}
