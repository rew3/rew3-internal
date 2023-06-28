package common

import (
	"strings"
	"time"

	"github.com/rew3/rew3-internal/service/common/request"
)

type FilterOperator string
type DateValueType string

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

	// Date Value Type.
	LAST_MONTH DateValueType = "last_month"
	LAST_WEEK  DateValueType = "last_week"
	YESTERDAY  DateValueType = "yesterday"
	TODAY      DateValueType = "today"
	TOMORROW   DateValueType = "tomorrow"
	NEXT_WEEK  DateValueType = "next_week"
	NEXT_MONTH DateValueType = "next_month"
)

type Filter struct {
	Field    string
	Value    interface{}
	Operator FilterOperator
}

func ParseFilterParams(param request.RequestParam) []Filter {
	filters := param.Filters
	splits := strings.Split(filters, "&")
	result := []Filter{}
	for _, f := range splits {
		filter := parseFilter(f)
		result = append(result, filter)
	}
	return result
}

func parseFilter(filter string) Filter {
	splits := strings.Split(filter, "=")
	if len(splits) == 2 {
		key := splits[0]
		value := splits[1]
		fieldName, op := parseFieldAndOperator(key)
		return Filter{
			Field:    fieldName,
			Value:    value,
			Operator: op,
		}
	} else {
		key := splits[0]
		fieldName, op := parseFieldAndOperator(key)
		return Filter{
			Field:    fieldName,
			Value:    nil,
			Operator: op,
		}
	}
}
// TODO parse, in, range etc. 
func parseDateValue() time.Time {

}

func parseFieldAndOperator(key string) (string, FilterOperator) {
	parts := strings.Split(key, "-")
	if len(parts) != 2 {
		// No filter operator type applied.
		return parts[0], NONE
	}
	fieldName := parts[0]
	operator := parts[1]
	switch operator {
	case string(EQ):
		return fieldName, EQ
	case string(NOT_EQ):
		return fieldName, EQ
	case string(EMPTY):
		return fieldName, EQ
	case string(NOT_EMPTY):
		return fieldName, EQ
	case string(CONTAINS):
		return fieldName, EQ
	case string(NOT_CONTAINS):
		return fieldName, EQ
	case string(STARTS_WITH):
		return fieldName, EQ
	case string(NOT_STARTS_WITH):
		return fieldName, EQ
	case string(ENDS_WITH):
		return fieldName, EQ
	case string(NOT_ENDS_WITH):
		return fieldName, EQ
	case string(LESS_THAN):
		return fieldName, EQ
	case string(NOT_LESS_THAN):
		return fieldName, EQ
	case string(LESS_THAN_EQUAL):
		return fieldName, EQ
	case string(NOT_LESS_THAN_EQUAL):
		return fieldName, EQ
	case string(GREATER_THAN):
		return fieldName, EQ
	case string(NOT_GREATER_THAN):
		return fieldName, EQ
	case string(GREATER_THAN_EQUAL):
		return fieldName, EQ
	case string(NOT_GREATER_THAN_EQUAL):
		return fieldName, EQ
	case string(IN):
		return fieldName, EQ
	case string(NOT_IN):
		return fieldName, EQ
	case string(RANGE):
		return fieldName, EQ
	case string(NOT_IN_RANGE):
		return fieldName, EQ
	default:
		return fieldName, NONE
	}
}
