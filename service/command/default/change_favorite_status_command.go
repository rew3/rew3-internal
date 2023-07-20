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
type ChangeFavoriteStatusCommandHandler[T common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[T]
}

type ChangeFavoriteStatusCommandHandlerContext[M common.Model, C command.Command] struct {
	CmdToData         func(C) (string, bool)
	FavoriteFieldName func() string
}

/**
 * Handle Command.
 */
func (ch *ChangeFavoriteStatusCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext ChangeFavoriteStatusCommandHandlerContext[M, C]) command.CommandResult {
	id, status := hContext.CmdToData(cmd)
	response, err := ch.changeStatus(ctx, id, status, hContext)
	return GenerateCmdResult[M](id, response, err, "Change"+ch.EntityName+"FavoriteStatus")
}

/**
 * Delete Record.
 */
func (ch *ChangeFavoriteStatusCommandHandler[M, C]) changeStatus(ctx context.Context, id string, status bool, hContext ChangeFavoriteStatusCommandHandlerContext[M, C]) (*M, error) {
	rc, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		contextUserId := rc.User.Id
		data := make(map[string]interface{})
		data["is_favourite"] = status
		if _, err := ch.Repository.AppendToArrayField(ctx, id, hContext.FavoriteFieldName(), "user._id", contextUserId, data); err != nil {
			return nil, err
		} else {
			return record, nil
		}
	} else {
		return nil, errors.New(ch.EntityName + " not found for given id")
	}
}
