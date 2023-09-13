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
 * Update command handler.
 * This handler can be used to update record for any entity/model type.
 */
type UpdateCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type UpdateCommandHandlerContext[M common.Model, C command.Command] struct {
	CmdToModel func(*C) (M, error)
}

/**
 * Handle Command.
 */
func (ch *UpdateCommandHandler[M, C]) Handle(ctx context.Context, id string, cmd C, hContext UpdateCommandHandlerContext[M, C]) command.CommandResult[M] {
	model, err := hContext.CmdToModel(&cmd)
	if ok, cmdResult := HandleError[M](err, "Update"+ch.EntityName); !ok {
		return cmdResult
	}
	response, status, err := ch.update(ctx, id, model)
	if err != nil {
		return command.CommandResult[M]{
			Response: r.ErrorExecutionResult[M]("Update"+ch.EntityName, err.Error(), status),
		}
	}
	return command.CommandResult[M]{
		Response: r.SuccessExecutionResult[M]("Update"+ch.EntityName, "Successfully record updated.", c.OK, *response),
	}
}

/**
 * Update Record.
 */
func (ch *UpdateCommandHandler[M, C]) update(ctx context.Context, id string, model M) (*M, c.StatusType, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, c.FORBIDDEN, errors.New("request context is not available")
	}
	update := &model
	if record := ch.Repository.FindById(ctx, id); record != nil {
		res, err := ch.Repository.Update(ctx, id, update)
		if err != nil {
			return nil, c.INTERNAL_SERVER_ERROR, err
		}
		return res, c.OK, nil
	} else {
		return nil, c.BAD_REQUEST, errors.New(ch.EntityName + " not found for given id")
	}
}
