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
type ChangeFavoriteStatusCommandHandler[W common.ModelWrapper, T common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[T]
	WrapperProvider func(*T) W
}

/**
 * Handle Command.
 */
func (ch *ChangeFavoriteStatusCommandHandler[W, T, C]) Handle(ctx context.Context, cmd C, resolveData func(C) (string, bool)) command.CommandResult {
	id, status := resolveData(cmd)
	response, err := ch.changeStatus(ctx, id, status)
	return GenerateCmdResult[T](id, *response, err, "Change"+ch.EntityName+"FavoriteStatus")
}

/**
 * Delete Record.
 */
func (ch *ChangeFavoriteStatusCommandHandler[W, T, C]) changeStatus(ctx context.Context, id string, status bool) (*T, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		data := ch.WrapperProvider(record)
		data.SetFavorite(status)
		if _, err := ch.Repository.Update(ctx, id, record); err != nil {
			return nil, err
		} else {
			return record, nil
		}
	} else {
		return nil, errors.New(ch.EntityName + " not found for given id")
	}
}
