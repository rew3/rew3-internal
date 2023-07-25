package server

import (
	"context"
	"encoding/json"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
)

/**
 * Service Method. 
 * Used to define grpc service method and handle it. 
 */
type ServiceMethod struct {
	api     api.APIOperation
	handler func(context.Context, request.RequestContext, json.RawMessage) *grpc.ResponsePayload
}

/**
 * New instance of Service Method.
 */
func NewServiceMethod(api api.APIOperation) *ServiceMethod {
	return &ServiceMethod{api: api}
}

/**
 * Bind handler for this service method. 
 * A handler is responsible for handling this service method execution and returning response as required.
 */
func (sm *ServiceMethod) BindHandler(handler func(context.Context, request.RequestContext, json.RawMessage) *grpc.ResponsePayload) *ServiceMethod {
	sm.handler = handler
	return sm
}

/**
 * Call this service method. 
 * Note: to be used internally by service.
 */
func (sm *ServiceMethod) call(ctx context.Context, request grpc.RequestPayload) *grpc.ResponsePayload {
	return sm.handler(ctx, request.Context, request.Data)
}
