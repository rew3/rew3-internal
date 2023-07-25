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
	method    api.APIOperation
	inputType interface{}
	handler   ServiceMethodHandler
}

func NewServiceMethod(api api.APIOperation, inputType interface{}, handler ServiceMethodHandler) *ServiceMethod {
	return &ServiceMethod{
		method:    api,
		inputType: inputType,
		handler:   handler,
	}
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
	return sm.handler.HandleMethod(ctx, request.Context, sm.inputType)
}

type ServiceMethodHandler interface {
	HandleMethod(ctx context.Context, rc request.RequestContext, data interface{}) *grpc.ResponsePayload
}
