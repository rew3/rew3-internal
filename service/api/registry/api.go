package server

import (
	"context"

	"github.com/rew3/rew3-internal/service/api/endpoints"
	grpc "github.com/rew3/rew3-internal/service/grpc/payload"
	"github.com/rew3/rew3-internal/service/shared/request"
	"github.com/rew3/rew3-internal/service/shared/response"
)

/**
 * Service API.
 * Used to define grpc service method and handle it.
 */
type ServiceAPI struct {
	api     endpoints.Endpoint
	handler func(context.Context, request.RequestContext, *response.ExecutionResult[interface{}]) *grpc.ResponsePayload
}

/**
 * New instance of Service Method.
 */
func NewServiceAPI(api endpoints.Endpoint) *ServiceAPI {
	return &ServiceAPI{api: api}
}

/**
 * Bind handler for this service method.
 * A handler is responsible for handling this service method execution and returning response as required.
 */
func (sm *ServiceAPI) BindHandler(handler func(context.Context, request.RequestContext, *response.ExecutionResult[interface{}]) *grpc.ResponsePayload) *ServiceAPI {
	sm.handler = handler
	return sm
}

/**
 * Call this service method.
 * Note: to be used internally by service.
 */
func (sm *ServiceAPI) call(ctx context.Context, request grpc.RequestPayload) *grpc.ResponsePayload {
	return sm.handler(ctx, request.Context, grpc.WrapPayload(request.Data))
}
