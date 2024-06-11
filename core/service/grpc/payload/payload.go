package payload

import (
	"encoding/json"

	"github.com/rew3/rew3-pkg/core/service/api"
	"github.com/rew3/rew3-pkg/core/service/shared/request"
	"github.com/rew3/rew3-pkg/core/service/shared/response/constants"
)

/**
 * Request and Response payloads.
 */
type RequestPayload struct {
	API     api.Endpoint
	Context request.RequestContext
	Data    json.RawMessage
}

type ResponsePayload struct {
	API           api.Endpoint
	Status        constants.StatusType
	StatusMessage string
	Data          interface{}
}

/**
 * Parse given execution result nto ResponsePayload.
 */
func ToResponsePayload(api api.Endpoint, status constants.StatusType, statusMessage string, data interface{}) *ResponsePayload {
	return &ResponsePayload{
		API:           api,
		Status:        status,
		StatusMessage: statusMessage,
		Data:          data,
	}
}

/**
 * Invalid Request Payload builder.
 */
func InvalidRequestResponsePayload(api api.Endpoint, err string) ResponsePayload {
	return ResponsePayload{
		API:           api,
		Status:        constants.BAD_REQUEST,
		StatusMessage: err,
	}
}
