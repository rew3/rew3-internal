package query

import (
	"fmt"
	"math/big"
	"time"
)

/**
 * Constants for Query Operator. 
 */
type QueryOperator string

const (
	// For All Data Types.
	EQUAL     QueryOperator = "equal"
	EMPTY     QueryOperator = "empty"

	// For String.
	MATCHES 	QueryOperator = "matches"
	STARTS_WITH QueryOperator = "starts_with"
	ENDS_WITH   QueryOperator = "ends_with"

	// For String and Numbers.
	LESS_THAN          	QueryOperator = "lt"
	LESS_THAN_EQUAL    	QueryOperator = "lte"
	GREATER_THAN       	QueryOperator = "gt"
	GREATER_THAN_EQUAL 	QueryOperator = "gte"
	IN                 	QueryOperator = "in"
	RANGE 				QueryOperator = "range"
)

/**
 * Root DSL for Query Definition. 
 */
type RootDSL struct {
	Queries []QueryDSL
}

/**
 * Interface for QUERY DSL. 
 * This is the base DSL type for all queries. 
 */
type QueryDSL interface{}

/*
 * DSL for defining criteria for particular field. 
 */
type CriteriaDSL struct {
	QueryDSL
	Field    FieldDSL
	Value    ValueType
	Operator QueryOperator
}

/*
 * DSL for defining logical operators. 
 */
type LogicalDSL struct {
	QueryDSL
}

/*
 * DSL for defining AND logical operator. 
 */
type ANDLogicalDSL struct {
	LogicalDSL
	Query []QueryDSL
}

/*
 * DSL for defining OR logical operator. 
 */
type ORLogicalDSL struct {
	LogicalDSL
	Query []QueryDSL
}

/*
 * DSL for defining NOT logical operator. 
 */
type NOTLogicalDSL struct {
	LogicalDSL
	Query []LogicalDSL
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
type NumberFieldDSL struct{ Field FieldDSL }
type DateTimeFieldDSL struct{ Field FieldDSL }

/*
 * Convert field to type safe dsl. 
 * Using this, only query methods related to particular data type is allowed. 
 */
 func (field FieldDSL) BoolType() BooleanFieldDSL	{ return BooleanFieldDSL{Field: field} }
 func (field FieldDSL) StrType() StringFieldDSL    	{ return StringFieldDSL{Field: field} }
 func (field FieldDSL) NumType() NumberFieldDSL    	{ return NumberFieldDSL{Field: field} }
 func (field FieldDSL) DateType() DateTimeFieldDSL 	{ return DateTimeFieldDSL{Field: field} }

 // Client Methods. 

/**
 * Root DSL for query definition. 
 */
func Query(dsl ...QueryDSL) RootDSL {
	return RootDSL{Queries: dsl}
}

/**
 * Logical AND query definition. 
 */
func AND(dsl ...QueryDSL) ANDLogicalDSL {
	return ANDLogicalDSL{Query: dsl}
}

/**
 * Logical OR query definition. 
 */
func OR(dsl ...QueryDSL) ORLogicalDSL {
	return ORLogicalDSL{Query: dsl}
}

/**
 * Logical NOT query definition. 
 */
func NOT(dsl ...LogicalDSL) NOTLogicalDSL {
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
func criteria(field FieldDSL, value interface{}, operator QueryOperator) CriteriaDSL {
	return CriteriaDSL{Field: field, Value: resolveScalarValue(value), Operator: operator}
}

/*
 * Build instance of Criteria DSL for In operator. 
 */
func criteriaForIn(field FieldDSL, value []interface{}) CriteriaDSL {
	resolvedValues := []ScalarValue{}
	for _, item := range value {
		resolvedValues = append(resolvedValues, resolveScalarValue(item))
	}
   return CriteriaDSL{Field: field, Value: ListValue{Value: resolvedValues}, Operator: IN}
}

/*
 * Build instance of Criteria DSL for Range Operator. 
 */
func criteriaForRange(field FieldDSL, start interface{}, end interface{}) CriteriaDSL {
   return CriteriaDSL{Field: field, Value: RangeValue{Start: resolveScalarValue(start), End: resolveScalarValue(end)}, Operator: RANGE}
}

// List of Generic Operators (Not Type Safe.)
func (field FieldDSL) Eq(value interface{}) CriteriaDSL { return criteria(field, value, EQUAL) }
func (field FieldDSL) Empty(value interface{}) CriteriaDSL { return criteria(field, nil, EMPTY) }
func (field FieldDSL) Matches(value interface{}) CriteriaDSL { return criteria(field, value, MATCHES) }
func (field FieldDSL) StartsWith(value interface{}) CriteriaDSL { return criteria(field, value, STARTS_WITH) }
func (field FieldDSL) EndsWith(value interface{}) CriteriaDSL { return criteria(field, value, ENDS_WITH) }
func (field FieldDSL) Lt(value interface{}) CriteriaDSL { return criteria(field, value, LESS_THAN) }
func (field FieldDSL) Lte(value interface{}) CriteriaDSL { return criteria(field, value, LESS_THAN_EQUAL) }
func (field FieldDSL) Gt(value interface{}) CriteriaDSL { return criteria(field, value, GREATER_THAN) }
func (field FieldDSL) Gte(value interface{}) CriteriaDSL { return criteria(field, value, GREATER_THAN_EQUAL) }
func (field FieldDSL) In(value ...interface{}) CriteriaDSL { return criteriaForIn(field, value) }
func (field FieldDSL) Range(start interface{}, end interface{}) CriteriaDSL { return criteriaForRange(field, start, end) }

// Type Safe operators. 

// Applies to Strings.
func (field StringFieldDSL) Eq(value interface{}) CriteriaDSL { return criteria(field.Field, value, EQUAL) }
func (field StringFieldDSL) Empty(value interface{}) CriteriaDSL { return criteria(field.Field, nil, EMPTY) }
func (field StringFieldDSL) Matches(value interface{}) CriteriaDSL { return criteria(field.Field, value, MATCHES) }
func (field StringFieldDSL) StartsWith(value interface{}) CriteriaDSL { return criteria(field.Field, value, STARTS_WITH) }
func (field StringFieldDSL) EndsWith(value interface{}) CriteriaDSL { return criteria(field.Field, value, ENDS_WITH) }
func (field StringFieldDSL) Lt(value interface{}) CriteriaDSL { return criteria(field.Field, value, LESS_THAN) }
func (field StringFieldDSL) Lte(value interface{}) CriteriaDSL { return criteria(field.Field, value, LESS_THAN_EQUAL) }
func (field StringFieldDSL) Gt(value interface{}) CriteriaDSL { return criteria(field.Field, value, GREATER_THAN) }
func (field StringFieldDSL) Gte(value interface{}) CriteriaDSL { return criteria(field.Field, value, GREATER_THAN_EQUAL) }
func (field StringFieldDSL) In(value ...interface{}) CriteriaDSL { return criteriaForIn(field.Field, value) }
func (field StringFieldDSL) Range(start interface{}, end interface{}) CriteriaDSL { return criteriaForRange(field.Field, start, end) }

// Applies to Boolean
func (field BooleanFieldDSL) Eq(value interface{}) CriteriaDSL { return criteria(field.Field, value, EQUAL) }
func (field BooleanFieldDSL) Empty(value interface{}) CriteriaDSL { return criteria(field.Field, nil, EMPTY) }

// Applies to Numbers
func (field NumberFieldDSL) Eq(value interface{}) CriteriaDSL { return criteria(field.Field, value, EQUAL) }
func (field NumberFieldDSL) Empty(value interface{}) CriteriaDSL { return criteria(field.Field, nil, EMPTY) }
func (field NumberFieldDSL) Lt(value interface{}) CriteriaDSL { return criteria(field.Field, value, LESS_THAN) }
func (field NumberFieldDSL) Lte(value interface{}) CriteriaDSL { return criteria(field.Field, value, LESS_THAN_EQUAL) }
func (field NumberFieldDSL) Gt(value interface{}) CriteriaDSL { return criteria(field.Field, value, GREATER_THAN) }
func (field NumberFieldDSL) Gte(value interface{}) CriteriaDSL { return criteria(field.Field, value, GREATER_THAN_EQUAL) }
func (field NumberFieldDSL) In(value ...interface{}) CriteriaDSL { return criteriaForIn(field.Field, value) }
func (field NumberFieldDSL) Range(start interface{}, end interface{}) CriteriaDSL { return criteriaForRange(field.Field, start, end) }

// Applies to DateTime
func (field DateTimeFieldDSL) Eq(value interface{}) CriteriaDSL { return criteria(field.Field, value, EQUAL) }
func (field DateTimeFieldDSL) Empty(value interface{}) CriteriaDSL { return criteria(field.Field, nil, EMPTY) }
func (field DateTimeFieldDSL) Lt(value interface{}) CriteriaDSL { return criteria(field.Field, value, LESS_THAN) }
func (field DateTimeFieldDSL) Lte(value interface{}) CriteriaDSL { return criteria(field.Field, value, LESS_THAN_EQUAL) }
func (field DateTimeFieldDSL) Gt(value interface{}) CriteriaDSL { return criteria(field.Field, value, GREATER_THAN) }
func (field DateTimeFieldDSL) Gte(value interface{}) CriteriaDSL { return criteria(field.Field, value, GREATER_THAN_EQUAL) }
func (field DateTimeFieldDSL) In(value ...interface{}) CriteriaDSL { return criteriaForIn(field.Field, value) }
func (field DateTimeFieldDSL) Range(start interface{}, end interface{}) CriteriaDSL { return criteriaForRange(field.Field, start, end) }

// TODO - Use elemMatch.
/*func (field FieldDSL) ArrayMatch(fieldDSLs ...FieldDSL) CriteriaDSL {...}*/

/*
 * Type Definition for Value Type. 
 */
type ValueType interface{}
type ScalarValue interface{
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
	Start 	ScalarValue
	End 	ScalarValue
}

/**
 * Resolve value for given generic type. 
 * Note: If given value is other than primitive types, value is stringified. 
 */
func resolveScalarValue(value interface{}) ScalarValue {
	switch v := value.(type) {
	case string:
		return StringValue{Value: v}
	case bool:
		return BooleanValue{Value: v}
	case int:
		return IntValue{Value: int32(v)}
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
