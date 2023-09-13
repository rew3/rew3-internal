package grpc

import (
	"encoding/json"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc/api"

	baseConst "github.com/rew3/rew3-internal/app/common/constants"
)

/**
 * Request and Response payloads.
 */
type RequestPayload struct {
	API     api.APIOperation
	Context request.RequestContext
	Data    json.RawMessage
}

type ResponsePayload struct {
	API           api.APIOperation
	Status        constants.StatusType
	StatusMessage string
	Data          json.RawMessage
	DataMeta      DataMeta
}

type DataMeta struct {
	Type interface{} // Type is among one of Empty, Binary, List, Object and Scalar.
}

type ScalarType string

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

type DataType interface {
	GetType() interface{}
}
type Empty struct{}

func (t Empty) GetType() interface{} {
	return t
}

type Binary struct{}

func (t Binary) GetType() interface{} {
	return t
}

type List struct {
	Type baseConst.Entity
}

func (t List) GetType() interface{} {
	return t
}

type Object struct {
	Type baseConst.Entity
}

func (t Object) GetType() interface{} {
	return t
}

type Scalar struct {
	Type ScalarType
}

func (t Scalar) GetType() interface{} {
	return t
}

/**
 * Parse given execution result nto ResponsePayload.
 */
/*func ToResponsePayload[T any](api api.APIOperation, executionResult *response.ExecutionResult[T]) *ResponsePayload {
	var data json.RawMessage = nil
	if parsed, err := ToJson[T](executionResult.Data); err == nil {
		data = parsed
	}
	return &ResponsePayload{
		API:           api,
		Status:        executionResult.Status,
		StatusMessage: executionResult.Message,
		Data:          data,
	}
}*/

/**
 * Parse given execution result nto ResponsePayload.
 */
func ToResponsePayload(api api.APIOperation, status constants.StatusType,
	statusMessage string, data interface{}, dataType DataType) *ResponsePayload {
	var rawdata json.RawMessage = nil
	if parsed, err := ToJson[interface{}](data); err == nil {
		rawdata = parsed
	}
	return &ResponsePayload{
		API:           api,
		Status:        status,
		StatusMessage: statusMessage,
		Data:          rawdata,
		DataMeta:      DataMeta{Type: dataType.GetType()},
	}
}
