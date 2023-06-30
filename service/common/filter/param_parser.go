package filter

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/rew3/rew3-internal/service/common/request"
)

/**
 * Filter parameter Parser utility.
 */

type FilterOperator string

const (
	NONE FilterOperator = ""

	// For All Data Types.
	EQ        FilterOperator = "-eq"
	NOT_EQ    FilterOperator = "-not-eq"
	EMPTY     FilterOperator = "-empty"
	NOT_EMPTY FilterOperator = "-not-empty"

	// For String.
	CONTAINS        FilterOperator = "-contains"
	NOT_CONTAINS    FilterOperator = "-not-contains"
	STARTS_WITH     FilterOperator = "-starts-with"
	NOT_STARTS_WITH FilterOperator = "-not-starts-with"
	ENDS_WITH       FilterOperator = "-ends-with"
	NOT_ENDS_WITH   FilterOperator = "-not-ends-with"

	// For String and Numbers.
	LESS_THAN              FilterOperator = "-lt"
	NOT_LESS_THAN          FilterOperator = "-not-lt"
	LESS_THAN_EQUAL        FilterOperator = "-lt-eq"
	NOT_LESS_THAN_EQUAL    FilterOperator = "-not-lt-eq"
	GREATER_THAN           FilterOperator = "-gt"
	NOT_GREATER_THAN       FilterOperator = "-not-gt"
	GREATER_THAN_EQUAL     FilterOperator = "-gt-eq"
	NOT_GREATER_THAN_EQUAL FilterOperator = "-not-gt-eq"
	IN                     FilterOperator = "-in"
	NOT_IN                 FilterOperator = "-not-in"
	RANGE                  FilterOperator = "-range"
	NOT_IN_RANGE           FilterOperator = "-not-in-range"
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
	if(param.Filters == "") {
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
		fieldName, op := parseFieldAndOperator(key)
		if op == RANGE || op == NOT_IN_RANGE {
			parsedValue, err := parseRangeValue(value)
			if err != nil {
				return Filter{}, fmt.Errorf("request parameter error. invalid range value `%s` is provided for parameter `%s`", value, key)
			}
			return Filter{
				Field:    fieldName,
				Value:    parsedValue,
				Operator: op,
			}, nil
		} else if op == IN || op == NOT_IN {
			parsedValue, err := parseListValue(value)
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
				Value:    parseValue(value),
				Operator: op,
			}, nil
		}
	} else {
		key := splits[0]
		fieldName, op := parseFieldAndOperator(key)
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
func parseFieldAndOperator(key string) (string, FilterOperator) {
	index := strings.Index(key, "-")
	if index < 0 {
		// No filter operator type applied.
		return key, NONE
	} else if key[:index] == "" {
		return key, NONE
	}
	fieldName := key[:index]
	operator := key[index:]
	switch operator {
	case string(EQ):
		return fieldName, EQ
	case string(NOT_EQ):
		return fieldName, NOT_EQ
	case string(EMPTY):
		return fieldName, EMPTY
	case string(NOT_EMPTY):
		return fieldName, NOT_EMPTY
	case string(CONTAINS):
		return fieldName, CONTAINS
	case string(NOT_CONTAINS):
		return fieldName, NOT_CONTAINS
	case string(STARTS_WITH):
		return fieldName, STARTS_WITH
	case string(NOT_STARTS_WITH):
		return fieldName, NOT_STARTS_WITH
	case string(ENDS_WITH):
		return fieldName, ENDS_WITH
	case string(NOT_ENDS_WITH):
		return fieldName, NOT_ENDS_WITH
	case string(LESS_THAN):
		return fieldName, LESS_THAN
	case string(NOT_LESS_THAN):
		return fieldName, NOT_LESS_THAN
	case string(LESS_THAN_EQUAL):
		return fieldName, LESS_THAN_EQUAL
	case string(NOT_LESS_THAN_EQUAL):
		return fieldName, NOT_LESS_THAN_EQUAL
	case string(GREATER_THAN):
		return fieldName, GREATER_THAN
	case string(NOT_GREATER_THAN):
		return fieldName, NOT_GREATER_THAN
	case string(GREATER_THAN_EQUAL):
		return fieldName, GREATER_THAN_EQUAL
	case string(NOT_GREATER_THAN_EQUAL):
		return fieldName, NOT_GREATER_THAN_EQUAL
	case string(IN):
		return fieldName, IN
	case string(NOT_IN):
		return fieldName, NOT_IN
	case string(RANGE):
		return fieldName, RANGE
	case string(NOT_IN_RANGE):
		return fieldName, NOT_IN_RANGE
	default:
		return fieldName, NONE
	}
}

/**
 * Parse Value to appropriate data type.
 */
func parseValue(value string) interface{} {
	if parsed, err := parseDateTimeValue(value); err == nil {
		return parsed
	} else {
		return parsePrimitiveValue(value)
	}
}

/**
 * Parse range value.
 */
func parseRangeValue(tupleStr string) ([2]interface{}, error) {
	tupleStr = strings.Trim(tupleStr, "[]")
	elements := strings.Split(tupleStr, ",")
	if len(elements) != 2 {
		return [2]interface{}{}, fmt.Errorf("invalid tuple format: expected 2 elements, got %d", len(elements))
	}
	var parsedElements [2]interface{}
	for i, element := range elements {
		element = strings.TrimSpace(element)
		parsedElements[i] = parsePrimitiveValue(element)
	}
	return parsedElements, nil
}

/**
 * Parse list value.
 */
func parseListValue(listStr string) ([]interface{}, error) {
	listStr = strings.Trim(listStr, "[]")
	items := strings.Split(listStr, ",")
	var parsedItems []interface{}
	for _, item := range items {
		item = strings.TrimSpace(item)
		parsedItems = append(parsedItems, parsePrimitiveValue(item))
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
 * Parse primitive value.
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

	// Assume it as string if no match is found
	return value
}
