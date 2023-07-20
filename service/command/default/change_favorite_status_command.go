package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	r "github.com/rew3/rew3-internal/service/common/response"
	c "github.com/rew3/rew3-internal/service/common/response/constants"
	"go.mongodb.org/mongo-driver/bson"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Delete command handler.
 * This handler can be used to delete record for any entity/model type.
 *
 * Note: In order to use this command:
 * model must have favorite fields with type []FavoriteInfo.
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
func (ch *ChangeFavoriteStatusCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext ChangeFavoriteStatusCommandHandlerContext[M, C]) command.CommandResult[M] {
	id, status := hContext.CmdToData(cmd)
	response, errStatus, err := ch.changeStatus(ctx, id, status, hContext)
	if err != nil {
		return command.CommandResult[M]{
			Response: r.ErrorExecutionResult[M](id, "Change"+ch.EntityName+"FavoriteStatus", err.Error(), errStatus),
		}
	}
	return command.CommandResult[M]{
		Response: r.SuccessExecutionResult[M](id, "Change"+ch.EntityName+"FavoriteStatus", "Successfully changed favorite status", c.OK, *response),
	}
}

/**
 * Delete Record.
 */
func (ch *ChangeFavoriteStatusCommandHandler[M, C]) changeStatus(ctx context.Context, id string, status bool, hContext ChangeFavoriteStatusCommandHandlerContext[M, C]) (*M, c.StatusType, error) {
	rc, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, c.UNAUTHORIZED, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		data := bson.D{
			{Key: "is_favourite", Value: true},
			{Key: "user", Value: bson.D{
				{Key: "_id", Value: rc.User.Id},
				{Key: "first_name", Value: rc.User.FirstName},
				{Key: "last_name", Value: rc.User.LastName},
			}},
		}
		if _, err := ch.Repository.AppendToArrayField(ctx, id, hContext.FavoriteFieldName(), "favorites", rc.User.Id, data); err != nil {
			return nil, c.INTERNAL_SERVER_ERROR, err
		} else {
			updatedRecord := ch.Repository.FindById(ctx, id)
			return updatedRecord, c.OK, nil
		}
	} else {
		return nil, c.BAD_REQUEST, errors.New(ch.EntityName + " not found for given id")
	}
}
