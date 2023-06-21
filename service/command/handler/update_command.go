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
type UpdateCommandHandler[T common.Model] struct {
	entityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *UpdateCommandHandler[T]) Handle(ctx context.Context,
	cmd command.Command,
	cmdToModel func(command.Command) (T, error),
	transformModel func(T) (T, error)) command.CommandResult {
	model, err := cmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "Update"+ch.entityName); !ok {
		return cmdResult
	}
	transformedModel, err := transformModel(model)
	if ok, transformResult := HandleError(err, "Update"+ch.entityName); !ok {
		return transformResult
	}
	response, err := ch.update(ctx, transformedModel)
	return GenerateCmdResult[T](*response, err, "Update"+ch.entityName)
}

/**
 * Update Record.
 */
func (ch *UpdateCommandHandler[T]) update(ctx context.Context, data T) (*T, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	id := data.GetId()
	if record := ch.Repository.FindById(ctx, id); record != nil {
		data.SetMeta((*record).GetMeta())
		return ch.Repository.Update(ctx, id, &data)
	} else {
		return nil, errors.New(ch.entityName + " not found for given id")
	}
}
