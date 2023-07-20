package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

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
func (ch *UpdateCommandHandler[M, C]) Handle(ctx context.Context, id string, cmd C, hContext UpdateCommandHandlerContext[M, C]) command.CommandResult {
	model, err := hContext.CmdToModel(&cmd)
	if ok, cmdResult := HandleError(err, "Update"+ch.EntityName); !ok {
		return cmdResult
	}
	id, response, err := ch.update(ctx, id, model)
	return GenerateCmdResult[M](id, response, err, "Update"+ch.EntityName)
}

/**
 * Update Record.
 */
func (ch *UpdateCommandHandler[M, C]) update(ctx context.Context, id string, model M) (string, *M, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return "", nil, errors.New("request context is not available")
	}
	update := &model
	if record := ch.Repository.FindById(ctx, id); record != nil {
		res, err := ch.Repository.Update(ctx, id, update)
		return id, res, err
	} else {
		return "", nil, errors.New(ch.EntityName + " not found for given id")
	}
}
