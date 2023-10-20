package server

import (
	"context"

	"github.com/rew3/rew3-internal/service/common/request"
	"github.com/rew3/rew3-internal/service/executor"
	"github.com/rew3/rew3-internal/service/grpc"
	"github.com/rew3/rew3-internal/service/grpc/api"

	ctxUtil "github.com/rew3/rew3-internal/pkg/context"
	baseCommand "github.com/rew3/rew3-internal/service/command"
	baseQuery "github.com/rew3/rew3-internal/service/query"
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
 * Map given API to its respective command executor.
 * Input - type for payload data to parse as input i.e. Command. Note: payload must be valid structure as its command.
 * Output - output response type returned by command handler.
 */
func MapCmdAPI[Input any, Output any](bc *MappingContext, api api.APIOperation, responseMeta grpc.DataType) {
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
 * Map given API to its respective query executor.
 * Input - type for payload data to parse as input i.e. Query. Note: payload must be valid structure as its query.
 * Output - output response type returned by query handler.
 */
func MapQueryAPI[Input any, Output any](bc *MappingContext, api api.APIOperation, responseMeta grpc.DataType) {
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
