package server

import (
	"encoding/json"
	"log"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
)

type RequestHandler struct {
	HandlerContext HandlerContext
}

type HandlerContext struct {
	RequestTypeMapping *RequestTypeMapping
	Receive            func(api.APIOperation, request.RequestContext, interface{}) grpc.ResponsePayload
}

func (handler *RequestHandler) Handle(request grpc.RequestPayload) grpc.ResponsePayload {
	apiRequestTypeMap := handler.HandlerContext.RequestTypeMapping.Mapping
	requestType, found := apiRequestTypeMap[request.API]
	if !found {
		return grpc.ResponsePayload{
			API:           request.API,
			Status:        constants.SERVICE_UNAVAILABLE,
			StatusMessage: string("INVALID API: " + request.API + " is not valid or is unavailable"),
		}
	}
	// Unmarshal the raw request data into the appropriate request type
	err := json.Unmarshal([]byte(request.Data), &requestType)
	if err != nil {
		log.Println("Failed to unmarshal request data: %v", err)
		return grpc.ResponsePayload{
			API:           request.API,
			Status:        constants.BAD_REQUEST,
			StatusMessage: string("INVALID REQUEST DATA: " + err.Error()),
		}
	}
	return handler.HandlerContext.Receive(request.API, request.Context, requestType)
}
