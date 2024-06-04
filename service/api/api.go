package api

import (
	"context"

	"github.com/rew3/rew3-internal/service/shared/request"
	req "github.com/rew3/rew3-internal/service/shared/request"
	res "github.com/rew3/rew3-internal/service/shared/response"
)

/**
 * Service API.
 * Used to define grpc service method and handle it.
 */
type ServiceAPI struct {
	api     Endpoint
	handler func(ctx context.Context, rctx req.RequestContext, data map[string]interface{}) res.ExecutionResult[interface{}]
}

/**
 * New instance of Service Method.
 */
func NewServiceAPI(api Endpoint) *ServiceAPI {
	return &ServiceAPI{api: api}
}

/**
 * Bind handler for this service method.
 * A handler is responsible for handling this service method execution and returning response as required.
 */
func (sm *ServiceAPI) BindHandler(handler func(context.Context, request.RequestContext, map[string]interface{}) res.ExecutionResult[interface{}]) *ServiceAPI {
	sm.handler = handler
	return sm
}

/**
 * Call this service method.
 * Note: to be used internally by service.
 */
func (sm *ServiceAPI) Call(ctx context.Context, rctx req.RequestContext, input map[string]interface{}) res.ExecutionResult[interface{}] {
	return sm.handler(ctx, rctx, input)
}
