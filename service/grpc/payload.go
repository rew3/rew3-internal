package grpc

import (
	"encoding/json"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc/api"
)

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
