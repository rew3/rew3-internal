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
 * Delete command handler.
 * This handler can be used to delete record for any entity/model type.
 */
type DeleteCommandHandler[T common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *DeleteCommandHandler[T, C]) Handle(ctx context.Context, id string) command.CommandResult {
	response, err := ch.delete(ctx, id)
	return GenerateCmdResult[T](*response, err, "Delete"+ch.EntityName)
}

/**
 * Delete Record.
 */
func (ch *DeleteCommandHandler[T, C]) delete(ctx context.Context, id string) (*T, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		if _, err := ch.Repository.Delete(ctx, id); err != nil {
			return nil, err
		} else {
			return record, nil
		}
	} else {
		return nil, errors.New(ch.EntityName + " not found for given id")
	}
}
