package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/common/response/constants"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
)

type ServiceMethod struct {
	api       api.APIOperation
	inputType interface{}
	handler   func(context.Context, request.RequestContext, interface{}) *grpc.ResponsePayload
}

func NewServiceMethod(api api.APIOperation) *ServiceMethod {
	return &ServiceMethod{api: api}
}

func (sm *ServiceMethod) call(ctx context.Context, request grpc.RequestPayload) *grpc.ResponsePayload {
	err := json.Unmarshal([]byte(request.Data), &sm.inputType)
	if err != nil {
		log.Println("failed to unmarshal request data: " + err.Error())
		return &grpc.ResponsePayload{
			API:           request.API,
			Status:        constants.BAD_REQUEST,
			StatusMessage: string("INVALID REQUEST DATA: " + err.Error()),
		}
	}
	return sm.handler(ctx, request.Context, sm.inputType)
}

func (sm *ServiceMethod) BindInput(inputType interface{}) *ServiceMethod {
	sm.inputType = inputType
	return sm
}

func (sm *ServiceMethod) BindHandler(handler func(context.Context, request.RequestContext, interface{}) *grpc.ResponsePayload) *ServiceMethod {
	sm.handler = handler
	return sm
}
