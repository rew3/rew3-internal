package api

import (
	"context"

	"github.com/rew3/rew3-internal/service/cqrs/executor"
	"github.com/rew3/rew3-internal/service/shared/request"
	"github.com/rew3/rew3-internal/service/shared/response/constants"

	baseCommand "github.com/rew3/rew3-internal/service/cqrs/command"
	baseQuery "github.com/rew3/rew3-internal/service/cqrs/query"
	res "github.com/rew3/rew3-internal/service/shared/response"
	ctxUtil "github.com/rew3/rew3-internal/service/utils/context"

	jsUtil "github.com/rew3/rew3-internal/pkg/utils/json"
)

/*
 * API Serivice Binder for CQRS based Service.
 * This maps the requried API to its respective command/query executor.
 */
 type CQRSAPIBinder struct {
	Registry        *APIRegistry
	CommandExecutor *executor.CommandExecutor
	QueryExecutor   *executor.QueryExecutor
}

// for service, non cqrs based. 
type APIBinder struct {
	Registry        *APIRegistry
}

/*
 * Add Command Service API - map given API to its respective command executor.
 * Input - type for payload data to parse as input i.e. Command. Note: payload must be valid structure as its command.
 * Output - output response type returned by command handler.
 */
func BindCommandAPI[Input any, Output any](bc *CQRSAPIBinder, enpoint Endpoint) {
	api := NewServiceAPI(enpoint).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input map[string]interface{}) res.ExecutionResult[interface{}] {
			parsed, err := jsUtil.MapToType[Input](input)
			if err != nil {
				return res.ErrorExecutionResult[interface{}](string(enpoint), "Invalid request data provided: "+err.Error(), constants.BAD_REQUEST)
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			var output Output
			cmdResult := bc.CommandExecutor.Execute(updatedCtx, parsed)
			response := baseCommand.ParseCommandResult[Output](&cmdResult, output)
			return  response.Generify()
		},
	)
	bc.Registry.AddServiceAPI(api)
}

/*
 * Add Query Service API - map given API to its respective query executor.
 * Input - type for payload data to parse as input i.e. Query. Note: payload must be valid structure as its query.
 * Output - output response type returned by query handler.
 */
func BindQueryAPI[Input any, Output any](bc *CQRSAPIBinder, enpoint Endpoint) {
	api := NewServiceAPI(enpoint).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input map[string]interface{}) res.ExecutionResult[interface{}] {
			parsed, err := jsUtil.MapToType[Input](input)
			if err != nil {
				return res.ErrorExecutionResult[interface{}](string(enpoint), "Invalid request data provided: "+err.Error(), constants.BAD_REQUEST)
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			var output Output
			queryResult := bc.QueryExecutor.Execute(updatedCtx, parsed)
			response := baseQuery.ParseQueryResult[Output](&queryResult, output)
			return  response.Generify()
		},
	)
	bc.Registry.AddServiceAPI(api)
}

/*
 * Bind service API for given endpoint with callback handler. 
 */
func BindAPI[Input any, Output any](binder *APIBinder, enpoint Endpoint, callback func(Input, context.Context)(res.ExecutionResult[Output])) {
	api := NewServiceAPI(enpoint).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input map[string]interface{}) res.ExecutionResult[interface{}] {
			parsed, err := jsUtil.MapToType[Input](input)
			if err != nil {
				return res.ErrorExecutionResult[interface{}](string(enpoint), "Invalid request data provided: "+err.Error(), constants.BAD_REQUEST)
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			result := callback(parsed, updatedCtx)
			return result.Generify()
		},
	)
	binder.Registry.AddServiceAPI(api)
}

