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
 * Delete command handler.
 * This handler can be used to delete record for any entity/model type.
 */
type DeleteCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

/**
 * Handle Command.
 */
func (ch *DeleteCommandHandler[M, C]) Handle(ctx context.Context, id string) command.CommandResult[M] {
	response, status, err := ch.delete(ctx, id)
	if err != nil {
		return command.CommandResult[M]{
			Response: r.ErrorExecutionResult[M]("Delete"+ch.EntityName, err.Error(), status),
		}
	}
	return command.CommandResult[M]{
		Response: r.SuccessExecutionResult[M]("Delete"+ch.EntityName, "Successfully record deleted.", c.OK, *response),
	}
}

/**
 * Delete Record.
 */
func (ch *DeleteCommandHandler[T, C]) delete(ctx context.Context, id string) (*T, c.StatusType, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, c.FORBIDDEN, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		if _, err := ch.Repository.Delete(ctx, id); err != nil {
			return nil, c.INTERNAL_SERVER_ERROR, err
		} else {
			return record, c.OK, nil
		}
	} else {
		return nil, c.BAD_REQUEST, errors.New(ch.EntityName + " not found for given id")
	}
}
