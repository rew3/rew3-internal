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
type DeleteCommandHandler[T common.Model] struct {
	entityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *DeleteCommandHandler[T]) Handle(ctx context.Context, id string) command.CommandResult {
	response, err := ch.delete(ctx, id)
	return GenerateCmdResult[T](*response, err, "Delete"+ch.entityName)
}

/**
 * Delete Record.
 */
func (ch *DeleteCommandHandler[T]) delete(ctx context.Context, id string) (*T, error) {
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
		return nil, errors.New(ch.entityName + " not found for given id")
	}
}
