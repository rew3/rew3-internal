package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	r "github.com/rew3/rew3-internal/service/common/response"
	c "github.com/rew3/rew3-internal/service/common/response/constants"

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
	CmdToModels func(*C) (map[string]*M, error)
}

/**
 * Handle Command.
 */
func (ch *BulkUpdateCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext BulkUpdateCommandHandlerContext[M, C]) command.CommandResult[interface{}] {
	models, err := hContext.CmdToModels(&cmd)
	if ok, cmdResult := HandleError[interface{}](err, "BulkUpdate"+ch.EntityName); !ok {
		return cmdResult
	}
	_, status, errs := ch.bulkUpdate(ctx, models)
	if errs != nil {
		return command.CommandResult[interface{}]{
			Response: r.ErrorExecutionResult[interface{}]("-", "BulkUpdate"+ch.EntityName, err.Error(), status),
		}
	}
	return command.CommandResult[interface{}]{
		Response: r.SuccessExecutionResult[interface{}]("-", "BulkUpdate"+ch.EntityName, "Successfully bulk updated records", c.OK, nil),
	}
}

/**
 * Bulk Update Record.
 */
func (ch *BulkUpdateCommandHandler[M, C]) bulkUpdate(ctx context.Context, data map[string]*M) (bool, c.StatusType, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, c.FORBIDDEN, errors.New("request context is not available")
	}
	_, err := ch.Repository.BulkUpdate(ctx, data)
	if err != nil {
		return false, c.INTERNAL_SERVER_ERROR, err
	}
	return true, c.OK, nil
}
