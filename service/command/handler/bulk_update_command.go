package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	bResponse "github.com/rew3/rew3-internal/service/common/response"
	bResponseConstant "github.com/rew3/rew3-internal/service/common/response/constants"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Bulk Update command handler.
 * This handler can be used to bulk update record for any entity/model type.
 */
type BulkUpdateCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type BulkUpdateCommandHandlerContext[M common.Model, C command.Command] struct {
	CmdToModels func(C) (map[string]*M, error)
}

/**
 * Handle Command.
 */
func (ch *BulkUpdateCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext BulkUpdateCommandHandlerContext[M, C]) command.CommandResult {
	models, err := hContext.CmdToModels(cmd)
	if ok, cmdResult := HandleError(err, "BulkUpdate"+ch.EntityName); !ok {
		return cmdResult
	}
	_, errs := ch.bulkUpdate(ctx, models)
	if errs != nil {
		return command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.INTERNAL_SERVER_ERROR,
				Message:      err.Error(),
				Action:       "BulkUpdate" + ch.EntityName,
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      ch.EntityName + " successfully bulk updated",
			Action:       "BulkUpdate" + ch.EntityName,
			Id:           "",
		},
	}
}

/**
 * Bulk Update Record.
 */
func (ch *BulkUpdateCommandHandler[M, C]) bulkUpdate(ctx context.Context, data map[string]*M) (bool, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	return ch.Repository.BulkUpdate(ctx, data)
}
