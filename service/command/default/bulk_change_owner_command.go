package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	r "github.com/rew3/rew3-internal/service/common/response"
	c "github.com/rew3/rew3-internal/service/common/response/constants"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * BulkChange owner command handler.
 * This handler can be used to bulk change owner of record for any entity/model type.
 */
type BulkChangeOwnerCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type BulkChangeOwnerCommandHandlerContext[M common.Model, C command.Command] struct {
	CmdToOwners func(*C) map[string]account.MiniUser
	SetOwner    func(*M, account.MiniUser)
}

/**
 * Handle Command.
 */
func (ch *BulkChangeOwnerCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext BulkChangeOwnerCommandHandlerContext[M, C]) command.CommandResult[interface{}] {
	owners := hContext.CmdToOwners(&cmd)
	_, status, err := ch.bulkChangeOwner(ctx, owners, hContext)
	if err != nil {
		return command.CommandResult[interface{}]{
			Response: r.ErrorExecutionResult[interface{}]("-", "BulkChange"+ch.EntityName+"Owner", err.Error(), status),
		}
	}
	return command.CommandResult[interface{}]{
		Response: r.SuccessExecutionResult[interface{}]("-", "BulkChange"+ch.EntityName+"Owner", "Successfully bulk changed record owners", c.OK, nil),
	}
}

/**
 * Bulk Change owner of Record.
 */
func (ch *BulkChangeOwnerCommandHandler[M, C]) bulkChangeOwner(
	ctx context.Context,
	data map[string]account.MiniUser,
	hContext BulkChangeOwnerCommandHandlerContext[M, C]) (bool, c.StatusType, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, c.FORBIDDEN, errors.New("request context is not available")
	}
	for id, owner := range data {
		if record := ch.Repository.FindById(ctx, id); record != nil {
			hContext.SetOwner(record, owner)
			ch.Repository.Update(ctx, id, record)
		}
	}
	return true, c.OK, nil
}
