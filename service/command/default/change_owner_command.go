package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Change owner command handler.
 * This handler can be used to change owner of record for any entity/model type.
 */
type ChangeOwnerCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type ChangeOwnerCommandHandlerContext[M common.Model, C command.Command] struct {
	ConvertCmd func(*C) (string, account.MiniUser)
	SetOwner   func(*M, account.MiniUser)
}

/**
 * Handle Command.
 */
func (ch *ChangeOwnerCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext ChangeOwnerCommandHandlerContext[M, C]) command.CommandResult {
	id, owner := hContext.ConvertCmd(&cmd)
	response, err := ch.changeOwner(ctx, id, owner, hContext)
	return GenerateCmdResult[M](id, response, err, "Change"+ch.EntityName+"Owner")
}

/**
 * Change owner of Record.
 */
func (ch *ChangeOwnerCommandHandler[M, C]) changeOwner(ctx context.Context, id string, owner account.MiniUser, hContext ChangeOwnerCommandHandlerContext[M, C]) (*M, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		hContext.SetOwner(record, owner)
		return ch.Repository.Update(ctx, id, record)
	} else {
		return nil, errors.New(ch.EntityName + " not found for given id")
	}
}
