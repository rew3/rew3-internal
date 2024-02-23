package shared

import (
	baseConst "github.com/rew3/rew3-internal/app/common/constants"
)

type Response struct {
	Type DataTypeEnum
	Data interface{}
}

/**
 * Contain data type information.
 */
type DataMeta struct {
	Type DataType // Type is among one of Empty, Binary, Object/Scalar List, Object and Scalar.
}

/*
 * Parent for all data types.
 */
type DataType interface {
	GetType() interface{}
}

/*
 * Empty type. Use this if response is not required or empty e.g. Null or Unit.
 */
type Empty struct{}

func (t Empty) GetType() interface{} {
	return t
}

/*
 * Binary Response type. Use this if response is binary.
 */
type Binary struct{}

func (t Binary) GetType() interface{} {
	return t
}

/*
 * List response type with data type.
 * Note: you can set Type to Nil, if you want to use entities other that defined in Entity.
 */
type List struct {
	Type baseConst.Entity
}

func (t List) GetType() interface{} {
	return t
}

/*
 * List response type with scalar data type. i.e. primitive type parameter.
 */
type ScalarList struct {
	Type ScalarType
}

func (t ScalarList) GetType() interface{} {
	return t
}

/*
 * Object response type.
 * Note: you can set Type to Nil, if you want to use entities other that defined in Entity.
 */
type Object struct {
	Type baseConst.Entity
}

func (t Object) GetType() interface{} {
	return t
}

/*
 * Scalar response type. e.g. primitive go types.
 */
type Scalar struct {
	Type ScalarType
}

func (t Scalar) GetType() interface{} {
	return t
}

/**
 * Scalar Type.
 */
type ScalarType string

type DataTypeEnum string

/**
 * List of scalar types / primitive types in golang.
 */
const (
	Int   ScalarType = "Int"
	Int8  ScalarType = "Int8"
	Int16 ScalarType = "Int16"
	Int32 ScalarType = "Int32"
	Int64 ScalarType = "Int64"

	UInt   ScalarType = "UInt"
	UInt8  ScalarType = "UInt8"
	UInt16 ScalarType = "UInt16"
	UInt32 ScalarType = "UInt32"
	UInt64 ScalarType = "UInt64"

	Float32 ScalarType = "Float32"
	Float64 ScalarType = "Float64"

	Complex64  ScalarType = "Complex64"
	Complex128 ScalarType = "Complex128"

	Bool   ScalarType = "Bool"
	String ScalarType = "String"
)
