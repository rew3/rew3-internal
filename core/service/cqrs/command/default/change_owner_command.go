package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-pkg/common/account"
	"github.com/rew3/rew3-pkg/core/service/cqrs/command"
	"github.com/rew3/rew3-pkg/core/service/shared"
	r "github.com/rew3/rew3-pkg/core/service/shared/response"
	c "github.com/rew3/rew3-pkg/core/service/shared/response/constants"
	"github.com/rew3/rew3-pkg/db/repository"

	rcUtil "github.com/rew3/rew3-pkg/core/service/utils/context"
)

/**
 * Change owner command handler.
 * This handler can be used to change owner of record for any entity/model type.
 */
type ChangeOwnerCommandHandler[M shared.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type ChangeOwnerCommandHandlerContext[M shared.Model, C command.Command] struct {
	ConvertCmd func(*C) (string, account.MiniUser)
	SetOwner   func(*M, account.MiniUser)
}

/**
 * Handle Command.
 */
func (ch *ChangeOwnerCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext ChangeOwnerCommandHandlerContext[M, C]) command.CommandResult[M] {
	id, owner := hContext.ConvertCmd(&cmd)
	response, status, err := ch.changeOwner(ctx, id, owner, hContext)
	if err != nil {
		return command.CommandResult[M]{
			Response: r.ErrorExecutionResult[M]("Change"+ch.EntityName+"Owner", err.Error(), status),
		}
	}
	return command.CommandResult[M]{
		Response: r.SuccessExecutionResult[M]("Change"+ch.EntityName+"Owner", "Successfully record deleted.", c.OK, *response),
	}

}

/**
 * Change owner of Record.
 */
func (ch *ChangeOwnerCommandHandler[M, C]) changeOwner(ctx context.Context, id string, owner account.MiniUser, hContext ChangeOwnerCommandHandlerContext[M, C]) (*M, c.StatusType, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, c.FORBIDDEN, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		hContext.SetOwner(record, owner)
		res, err := ch.Repository.Update(ctx, id, record)
		if err != nil {
			return nil, c.INTERNAL_SERVER_ERROR, err
		}
		return res, c.OK, nil
	} else {
		return nil, c.BAD_REQUEST, errors.New(ch.EntityName + " not found for given id")
	}
}
