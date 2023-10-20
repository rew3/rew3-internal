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

type ExecutorType string

const (
	COMMAND ExecutorType = "Command"
	QUERY   ExecutorType = "Query"
)

type MappingContext struct {
	registry        *ServiceMethodRegistry
	CommandExecutor *executor.CommandExecutor
	QueryExecutor   *executor.QueryExecutor
}

func MapAPI[Input any, Output any](bc *MappingContext, api api.APIOperation, exType ExecutorType, output Output, payloadResponseType grpc.DataType) {
	method := NewServiceMethod(api).BindHandler(
		func(ctx context.Context, rc request.RequestContext, input *grpc.PayloadWrapper) *grpc.ResponsePayload {
			parsed, err := grpc.ParseFullPayload[Input](input)
			if err != nil {
				return grpc.InvalidRequestResponsePayload(api, "Invalid request data provided: "+err.Error())
			}
			updatedCtx := ctxUtil.WithRequestContext(ctx, rc)
			if exType == COMMAND {
				cmdResult := bc.CommandExecutor.Execute(updatedCtx, parsed)
				er := baseCommand.ParseCommandResult[Output](&cmdResult, output)
				return grpc.ToResponsePayload(api, er.Status, er.Message, er.Data, payloadResponseType)
			} else {
				queryResult := bc.QueryExecutor.Execute(updatedCtx, parsed)
				er := baseQuery.ParseQueryResult[Output](&queryResult, output)
				return grpc.ToResponsePayload(api, er.Status, er.Message, er.Data, payloadResponseType)
			}
		},
	)
	bc.registry.AddServiceMethod(method)
}
