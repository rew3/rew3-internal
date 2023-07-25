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
}

/**
 * Parse given execution result nto ResponsePayload.
 */
func ToResponsePayload[T any](api api.APIOperation, executionResult response.ExecutionResult[T]) ResponsePayload {
	var data json.RawMessage = nil
	if parsed, err := ToJson[T](executionResult.Data); err == nil {
		data = parsed
	}
	return ResponsePayload{
		API:           api,
		Status:        executionResult.Status,
		StatusMessage: executionResult.Message,
		Data:          data,
	}
}
