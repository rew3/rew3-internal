package grpc

import (
	"encoding/json"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc/api"
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

type Empty struct{}
type Binary struct{}
type List struct {
	ParamType string
}
type Object struct {
	ObjectType string
}
type Scalar struct {
	ScalarType ScalarType
}

/**
 * Parse given execution result nto ResponsePayload.
 */
func ToResponsePayload[T any](api api.APIOperation, executionResult *response.ExecutionResult[T]) *ResponsePayload {
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
}
