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
type UpdateCommandHandler[W common.ModelWrapper, T common.Model, C command.Command] struct {
	EntityName      string
	Repository      repository.Repository[T]
	WrapperProvider func(*T) W
}

/**
 * Handle Command.
 */
func (ch *UpdateCommandHandler[W, T, C]) Handle(ctx context.Context,
	cmd C,
	cmdToModel func(C) (T, error),
	transformModel func(T) (T, error)) command.CommandResult {
	model, err := cmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "Update"+ch.EntityName); !ok {
		return cmdResult
	}
	transformedModel, err := transformModel(model)
	if ok, transformResult := HandleError(err, "Update"+ch.EntityName); !ok {
		return transformResult
	}
	id, response, err := ch.update(ctx, transformedModel)
	return GenerateCmdResult[T](id, response, err, "Update"+ch.EntityName)
}

/**
 * Update Record.
 */
func (ch *UpdateCommandHandler[W, T, C]) update(ctx context.Context, model T) (string, *T, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return "", nil, errors.New("request context is not available")
	}
	update := &model
	data := ch.WrapperProvider(update)
	id := data.GetId()
	if record := ch.Repository.FindById(ctx, id); record != nil {
		res, err := ch.Repository.Update(ctx, id, update)
		return id, res, err
	} else {
		return "", nil, errors.New(ch.EntityName + " not found for given id")
	}
}
