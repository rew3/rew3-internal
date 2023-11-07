package query

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

/**
 * Query DSL to define Query.
 *
 * Usage Example:
 * 	Query(
 *		Field("first_name").Eq("ranjit"),
 *		Field("last_name").NOT().Eq("kaliraj"),
 *		AND(
 *			Field("title").Eq("developer"),
 *			OR(
 *				Field("department").Eq("IT"),
 *				Field("level").Eq("senior"),
 *			),
 *		),
 *	)
 *
 */

/**
 * Constants for Query Operator.
 */
type QueryOperator string

const (
	UNKNOWN = "unknown"

	// For All Data Types.
	EQUAL QueryOperator = "equal"
	EMPTY QueryOperator = "empty"

	// For String.
	MATCHES     QueryOperator = "matches"
	STARTS_WITH QueryOperator = "starts_with"
	ENDS_WITH   QueryOperator = "ends_with"

	// For String and Numbers.
	LESS_THAN          QueryOperator = "lt"
	LESS_THAN_EQUAL    QueryOperator = "lte"
	GREATER_THAN       QueryOperator = "gt"
	GREATER_THAN_EQUAL QueryOperator = "gte"
	IN                 QueryOperator = "in"
	RANGE              QueryOperator = "range"
)

/**
 * Interface for QUERY DSL.
 * This is the base DSL type for all queries.
 */
type BaseDSL interface{}
type NonAssocDSL interface{}

/**
 * Root DSL for Query Definition.
 */
type QueryDSL struct {
	Queries []BaseDSL
}

/**
 * Associative Query DSL.
 * It represent a query in associative array.
 * e.g. {"first_name": "value", "last_name": "value"}
 */
type AssocDSL struct {
	BaseDSL
	Queries []NonAssocDSL
}

/*
 * DSL for defining logical operators.
 * Current supported Logical operator are AND, OR, NOT.
 */
type LogicalDSL struct {
	BaseDSL
	NonAssocDSL
}

/*
 * DSL for defining AND logical operator.
 */
type ANDLogicalDSL struct {
	LogicalDSL
	Queries []BaseDSL
}

/*
 * DSL for defining OR logical operator.
 */
type ORLogicalDSL struct {
	LogicalDSL
	Queries []BaseDSL
}

/*
 * DSL for defining NOT logical operator.
 */
type NOTLogicalDSL struct {
	LogicalDSL
	Query LogicalDSL
}

/*
 * DSL for defining criteria for particular field.
 */
type CriteriaDSL struct {
	BaseDSL
	NonAssocDSL
	Field    FieldDSL
	Value    ValueType
	Operator QueryOperator
}

/*
 * DSL for defining Field.
 */
type FieldDSL struct {
	Name       string
	IsNegation bool
}

/*
 * DSL for defining type specific field.
 * This DSL is used to access query operators specific to data type.
 */
type StringFieldDSL struct{ Field FieldDSL }
type BooleanFieldDSL struct{ Field FieldDSL }
type IntFieldDSL struct{ Field FieldDSL }
type Int32FieldDSL struct{ Field FieldDSL }
type Int64FieldDSL struct{ Field FieldDSL }
type Float32FieldDSL struct{ Field FieldDSL }
type Float64FieldDSL struct{ Field FieldDSL }
type BigIntFieldDSL struct{ Field FieldDSL }
type DateTimeFieldDSL struct{ Field FieldDSL }

/*
 * Convert field to type safe dsl.
 * Using this, only query methods related to particular data type is allowed.
 */
func (field FieldDSL) BoolType() BooleanFieldDSL    { return BooleanFieldDSL{Field: field} }
func (field FieldDSL) StrType() StringFieldDSL      { return StringFieldDSL{Field: field} }
func (field FieldDSL) IntType() IntFieldDSL         { return IntFieldDSL{Field: field} }
func (field FieldDSL) Int32Type() Int32FieldDSL     { return Int32FieldDSL{Field: field} }
func (field FieldDSL) Int64Type() Int64FieldDSL     { return Int64FieldDSL{Field: field} }
func (field FieldDSL) Float32Type() Float32FieldDSL { return Float32FieldDSL{Field: field} }
func (field FieldDSL) FLoat64Type() Float64FieldDSL { return Float64FieldDSL{Field: field} }
func (field FieldDSL) BigIntType() BigIntFieldDSL   { return BigIntFieldDSL{Field: field} }
func (field FieldDSL) DateType() DateTimeFieldDSL   { return DateTimeFieldDSL{Field: field} }

// Client Methods.

/**
 * Root DSL for query definition.
 */
func Query(dsl ...BaseDSL) QueryDSL {
	return QueryDSL{Queries: dsl}
}

/**
 * DSL for associative representation query definition.
 * This query DSL is used to define multiple query dsl into single dsl.
 * e.g. you can consider a json object: {key1: value1, key2: value2} where multiple key values
 * are defined in single object notation. This works the same way. 
 */
func Assoc(dsl ...NonAssocDSL) AssocDSL {
	return AssocDSL{Queries: dsl}
}

/**
 * Logical AND query definition.
 */
func AND(dsl ...BaseDSL) ANDLogicalDSL {
	return ANDLogicalDSL{Queries: dsl}
}

/**
 * Logical OR query definition.
 */
func OR(dsl ...BaseDSL) ORLogicalDSL {
	return ORLogicalDSL{Queries: dsl}
}

/**
 * Logical NOT query definition.
 */
func NOT(dsl LogicalDSL) NOTLogicalDSL {
	return NOTLogicalDSL{Query: dsl}
}

/**
 * Field Definition for given field name.
 */
func Field(name string) FieldDSL {
	return FieldDSL{Name: name}
}

/**
 * Define Negation in Field DSL.
 */
func (field FieldDSL) NOT() FieldDSL {
	f := &field
	f.IsNegation = true
	return *f
}

/*
 * Build instance of Criteria DSL.
 */
func Criteria(field FieldDSL, value interface{}, operator QueryOperator) CriteriaDSL {
	return CriteriaDSL{Field: field, Value: ToScalarValue(value), Operator: operator}
}

/*
 * Build instance of Criteria DSL for In operator.
 */
func CriteriaForIn(field FieldDSL, value []interface{}) CriteriaDSL {
	resolvedValues := []ScalarValue{}
	for _, item := range value {
		resolvedValues = append(resolvedValues, ToScalarValue(item))
	}
	return CriteriaDSL{Field: field, Value: ListValue{Value: resolvedValues}, Operator: IN}
}

/*
 * Build instance of Criteria DSL for Range Operator.
 */
func CriteriaForRange(field FieldDSL, start interface{}, end interface{}) CriteriaDSL {
	return CriteriaDSL{Field: field, Value: RangeValue{Start: ToScalarValue(start), End: ToScalarValue(end)}, Operator: RANGE}
}

// List of Generic Operators (Not Type Safe.)
func (field FieldDSL) IS(value interface{}) CriteriaDSL {
	return Criteria(field, value, UNKNOWN)
}
func (field FieldDSL) Eq(value interface{}) CriteriaDSL {
	return Criteria(field, value, EQUAL)
}
func (field FieldDSL) Empty() CriteriaDSL {
	return Criteria(field, nil, EMPTY)
}
func (field FieldDSL) Matches(value interface{}) CriteriaDSL {
	return Criteria(field, value, MATCHES)
}
func (field FieldDSL) StartsWith(value interface{}) CriteriaDSL {
	return Criteria(field, value, STARTS_WITH)
}
func (field FieldDSL) EndsWith(value interface{}) CriteriaDSL {
	return Criteria(field, value, ENDS_WITH)
}
func (field FieldDSL) Lt(value interface{}) CriteriaDSL {
	return Criteria(field, value, LESS_THAN)
}
func (field FieldDSL) Lte(value interface{}) CriteriaDSL {
	return Criteria(field, value, LESS_THAN_EQUAL)
}
func (field FieldDSL) Gt(value interface{}) CriteriaDSL {
	return Criteria(field, value, GREATER_THAN)
}
func (field FieldDSL) Gte(value interface{}) CriteriaDSL {
	return Criteria(field, value, GREATER_THAN_EQUAL)
}
func (field FieldDSL) In(value ...interface{}) CriteriaDSL {
	return CriteriaForIn(field, value)
}
func (field FieldDSL) Range(start interface{}, end interface{}) CriteriaDSL {
	return CriteriaForRange(field, start, end)
}

// Type Safe Query Operators.

// Applies to Strings.
func (field StringFieldDSL) IS(value string) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field StringFieldDSL) Eq(value string) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field StringFieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field StringFieldDSL) Matches(value string) CriteriaDSL {
	return Criteria(field.Field, value, MATCHES)
}
func (field StringFieldDSL) StartsWith(value string) CriteriaDSL {
	return Criteria(field.Field, value, STARTS_WITH)
}
func (field StringFieldDSL) EndsWith(value string) CriteriaDSL {
	return Criteria(field.Field, value, ENDS_WITH)
}
func (field StringFieldDSL) Lt(value string) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field StringFieldDSL) Lte(value string) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field StringFieldDSL) Gt(value string) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field StringFieldDSL) Gte(value string) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field StringFieldDSL) In(value ...string) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field StringFieldDSL) Range(start string, end string) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to Boolean
func (field BooleanFieldDSL) IS(value string) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field BooleanFieldDSL) Eq(value bool) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field BooleanFieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}

// Applies to Int
func (field IntFieldDSL) IS(value int) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field IntFieldDSL) Eq(value int) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field IntFieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field IntFieldDSL) Lt(value int) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field IntFieldDSL) Lte(value int) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field IntFieldDSL) Gt(value int) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field IntFieldDSL) Gte(value int) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field IntFieldDSL) In(value ...int) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field IntFieldDSL) Range(start int, end int) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to Int32
func (field Int32FieldDSL) IS(value int32) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field Int32FieldDSL) Eq(value int32) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field Int32FieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field Int32FieldDSL) Lt(value int32) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field Int32FieldDSL) Lte(value int32) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field Int32FieldDSL) Gt(value int32) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field Int32FieldDSL) Gte(value int32) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field Int32FieldDSL) In(value ...int32) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field Int32FieldDSL) Range(start int32, end int32) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to Int64
func (field Int64FieldDSL) IS(value int64) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field Int64FieldDSL) Eq(value int64) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field Int64FieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field Int64FieldDSL) Lt(value int64) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field Int64FieldDSL) Lte(value int64) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field Int64FieldDSL) Gt(value int64) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field Int64FieldDSL) Gte(value int64) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field Int64FieldDSL) In(value ...int64) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field Int64FieldDSL) Range(start int64, end int64) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to float32
func (field Float32FieldDSL) IS(value float32) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field Float32FieldDSL) Eq(value float32) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field Float32FieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field Float32FieldDSL) Lt(value float32) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field Float32FieldDSL) Lte(value float32) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field Float32FieldDSL) Gt(value float32) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field Float32FieldDSL) Gte(value float32) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field Float32FieldDSL) In(value ...float32) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field Float32FieldDSL) Range(start float32, end float32) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to float64
func (field Float64FieldDSL) IS(value float64) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field Float64FieldDSL) Eq(value float64) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field Float64FieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field Float64FieldDSL) Lt(value float64) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field Float64FieldDSL) Lte(value float64) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field Float64FieldDSL) Gt(value float64) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field Float64FieldDSL) Gte(value float64) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field Float64FieldDSL) In(value ...float64) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field Float64FieldDSL) Range(start float64, end float64) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to big.Int
func (field BigIntFieldDSL) IS(value big.Int) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field BigIntFieldDSL) Eq(value big.Int) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field BigIntFieldDSL) Empty() CriteriaDSL {
	return Criteria(field.Field, nil, EMPTY)
}
func (field BigIntFieldDSL) Lt(value big.Int) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field BigIntFieldDSL) Lte(value big.Int) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field BigIntFieldDSL) Gt(value big.Int) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field BigIntFieldDSL) Gte(value big.Int) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field BigIntFieldDSL) In(value ...big.Int) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field BigIntFieldDSL) Range(start big.Int, end big.Int) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// Applies to DateTime
func (field DateTimeFieldDSL) IS(value time.Time) CriteriaDSL {
	return Criteria(field.Field, value, UNKNOWN)
}
func (field DateTimeFieldDSL) Eq(value time.Time) CriteriaDSL {
	return Criteria(field.Field, value, EQUAL)
}
func (field DateTimeFieldDSL) Empty() CriteriaDSL { return Criteria(field.Field, nil, EMPTY) }
func (field DateTimeFieldDSL) Lt(value time.Time) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN)
}
func (field DateTimeFieldDSL) Lte(value time.Time) CriteriaDSL {
	return Criteria(field.Field, value, LESS_THAN_EQUAL)
}
func (field DateTimeFieldDSL) Gt(value time.Time) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN)
}
func (field DateTimeFieldDSL) Gte(value time.Time) CriteriaDSL {
	return Criteria(field.Field, value, GREATER_THAN_EQUAL)
}
func (field DateTimeFieldDSL) In(value ...time.Time) CriteriaDSL {
	return CriteriaForIn(field.Field, make([]interface{}, len(value)))
}
func (field DateTimeFieldDSL) Range(start time.Time, end time.Time) CriteriaDSL {
	return CriteriaForRange(field.Field, start, end)
}

// TODO - Use elemMatch.
/*func (field FieldDSL) ArrayMatch(fieldDSLs ...FieldDSL) CriteriaDSL {...}*/

/*
 * Type Definition for Value Type.
 */
type ValueType interface{}
type ScalarValue interface {
	ValueType
}
type StringValue struct {
	ScalarValue
	Value string
}
type BooleanValue struct {
	ScalarValue
	Value bool
}
type IntValue struct {
	ScalarValue
	Value int32
}
type LongValue struct {
	ScalarValue
	Value int64
}
type FloatValue struct {
	ScalarValue
	Value float32
}
type DoubleValue struct {
	ScalarValue
	Value float64
}
type BigIntValue struct {
	ScalarValue
	Value big.Int
}
type DateTimeValue struct {
	ScalarValue
	Value time.Time
}
type ListValue struct {
	ValueType
	Value []ScalarValue
}
type RangeValue struct {
	ValueType
	Start ScalarValue
	End   ScalarValue
}

/**
 * Resolve Scalar value from given value type.
 * Any other types than scalar will return nil value.
 */
func ResolveScalarValue(tpe ValueType) interface{} {
	switch v := tpe.(type) {
	case StringValue:
		return v.Value
	case BooleanValue:
		return v.Value
	case IntValue:
		return v.Value
	case LongValue:
		return v.Value
	case FloatValue:
		return v.Value
	case DoubleValue:
		return v.Value
	case BigIntValue:
		return v.Value
	case DateTimeValue:
		return v.Value
	default:
		return nil
	}
}

/**
 * Check if given Value is Scalar Value type or not.
 */
func IsScalarValue(tpe ValueType) bool {
	switch tpe.(type) {
	case StringValue:
		return true
	case BooleanValue:
		return true
	case IntValue:
		return true
	case LongValue:
		return true
	case FloatValue:
		return true
	case DoubleValue:
		return true
	case BigIntValue:
		return true
	case DateTimeValue:
		return true
	default:
		return false
	}
}

/**
 * Check if given Value is List Value type or not.
 */
func IsListValue(tpe ValueType) bool {
	switch tpe.(type) {
	case ListValue:
		return true
	default:
		return false
	}
}

/**
 * Check if given Value is Range Value type or not.
 */
func IsRangeValue(tpe ValueType) bool {
	switch tpe.(type) {
	case RangeValue:
		return true
	default:
		return false
	}
}

/**
 * Resolve value for given generic type.
 * Note: If given value is other than primitive types, value is stringified.
 */
func ToScalarValue(value interface{}) ScalarValue {
	switch v := value.(type) {
	case string:
		return StringValue{Value: v}
	case bool:
		return BooleanValue{Value: v}
	case int:
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return IntValue{Value: int32(v)}
		}
		return LongValue{Value: int64(v)}
	case int32:
		return IntValue{Value: v}
	case int64:
		return LongValue{Value: v}
	case float32:
		return FloatValue{Value: v}
	case float64:
		return DoubleValue{Value: v}
	case big.Int:
		return BigIntValue{Value: v}
	case time.Time:
		return DateTimeValue{Value: v}
	default:
		return StringValue{Value: fmt.Sprintf("%v", value)}
	}
}
