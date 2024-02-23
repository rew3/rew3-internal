package server

import (
	"context"

	"github.com/rew3/rew3-internal/service/cqrs/executor"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"
	"github.com/rew3/rew3-internal/service/shared/request"

	baseCommand "github.com/rew3/rew3-internal/service/cqrs/command"
	baseQuery "github.com/rew3/rew3-internal/service/cqrs/query"
	s "github.com/rew3/rew3-internal/service/shared/response"
	ctxUtil "github.com/rew3/rew3-internal/service/utils/context"
)

/*
 * API Serivice Mapping.
 * This utility maps the requried API to its respective command/query executor execution and register to GRPC Service method registry.
 */

type MappingContext struct {
	Registry        *ServiceMethodRegistry
	CommandExecutor *executor.CommandExecutor
	QueryExecutor   *executor.QueryExecutor
}

/*
 * Add Command Service - map given API to its respective command executor.
 * Input - type for payload data to parse as input i.e. Command. Note: payload must be valid structure as its command.
 * Output - output response type returned by command handler.
 */
func AddCommandService[Input any, Output any](bc *MappingContext, api api.APIOperation, responseMeta grpc.DataType) {
	method := NewServiceMethod(api).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input *grpc.PayloadWrapper) *grpc.ResponsePayload {
			parsed, err := grpc.ParseFullPayload[Input](input)
			if err != nil {
				return grpc.InvalidRequestResponsePayload(api, "Invalid request data provided: "+err.Error())
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			var output Output
			cmdResult := bc.CommandExecutor.Execute(updatedCtx, parsed)
			er := baseCommand.ParseCommandResult[Output](&cmdResult, output)
			return grpc.ToResponsePayload(api, er.Status, er.Message, er.Data, responseMeta)
		},
	)
	bc.Registry.AddServiceMethod(method)
}

/*
 * Add Query Service - map given API to its respective query executor.
 * Input - type for payload data to parse as input i.e. Query. Note: payload must be valid structure as its query.
 * Output - output response type returned by query handler.
 */
func AddQueryService[Input any, Output any](bc *MappingContext, api api.APIOperation, responseMeta grpc.DataType) {
	method := NewServiceMethod(api).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input *grpc.PayloadWrapper) *grpc.ResponsePayload {
			parsed, err := grpc.ParseFullPayload[Input](input)
			if err != nil {
				return grpc.InvalidRequestResponsePayload(api, "Invalid request data provided: "+err.Error())
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			var output Output
			queryResult := bc.QueryExecutor.Execute(updatedCtx, parsed)
			er := baseQuery.ParseQueryResult[Output](&queryResult, output)
			return grpc.ToResponsePayload(api, er.Status, er.Message, er.Data, responseMeta)
		},
	)
	bc.Registry.AddServiceMethod(method)
}

/*
 * Create new Service - with mapping of given API to its respective service handler.
 * You can use created service method and add to registry.
 * Output - output response type returned by query handler.
 */
func CreateService[Output any](api api.APIOperation, callback func(interface{}, context.Context) *s.ExecutionResult[Output], responseMeta grpc.DataType) *ServiceMethod {
	return NewServiceMethod(api).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input *grpc.PayloadWrapper) *grpc.ResponsePayload {
			parsed, err := grpc.ParseFullPayload[interface{}](input)
			if err != nil {
				return grpc.InvalidRequestResponsePayload(api, "Invalid request data provided: "+err.Error())
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			er := callback(parsed, updatedCtx)
			return grpc.ToResponsePayload(api, er.Status, er.Message, er.Data, responseMeta)
		},
	)
}
