package grpc

import (
	"encoding/json"

	"github.com/rew3/rew3-internal/service/api/endpoints"
	s "github.com/rew3/rew3-internal/service/shared"
	"github.com/rew3/rew3-internal/service/shared/request"
	"github.com/rew3/rew3-internal/service/shared/response/constants"
)

/**
 * Request and Response payloads.
 */
type RequestPayload struct {
	API     endpoints.Endpoint
	Context request.RequestContext
	Data    json.RawMessage
}

type ResponsePayload struct {
	API           endpoints.Endpoint
	Status        constants.StatusType
	StatusMessage string
	Data          interface{}
	DataMeta      s.DataMeta
}

/**
 * Parse given execution result nto ResponsePayload.
 */
func ToResponsePayload(api endpoints.Endpoint, status constants.StatusType,
	statusMessage string, data interface{}, dataType s.DataType) *ResponsePayload {
	return &ResponsePayload{
		API:           api,
		Status:        status,
		StatusMessage: statusMessage,
		Data:          data,
		DataMeta:      s.DataMeta{Type: dataType},
	}
}

/**
 * Invalid Request Payload builder.
 */
func InvalidRequestResponsePayload(api endpoints.Endpoint, err string) *ResponsePayload {
	return &ResponsePayload{
		API:           api,
		Status:        constants.BAD_REQUEST,
		StatusMessage: err,
	}
}
