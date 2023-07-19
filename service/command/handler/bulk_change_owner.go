package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	bResponse "github.com/rew3/rew3-internal/service/common/response"
	bResponseConstant "github.com/rew3/rew3-internal/service/common/response/constants"

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
	CmdToOwners func(C) map[string]account.MiniUser
	SetOwner    func(*M, account.MiniUser)
}

/**
 * Handle Command.
 */
func (ch *BulkChangeOwnerCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext BulkChangeOwnerCommandHandlerContext[M, C]) command.CommandResult {
	owners := hContext.CmdToOwners(cmd)
	_, err := ch.bulkChangeOwner(ctx, owners, hContext)
	if err != nil {
		return command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.INTERNAL_SERVER_ERROR,
				Message:      err.Error(),
				Action:       "BulkChange" + ch.EntityName + "Owner",
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      ch.EntityName + " successfully bulk owner changed",
			Action:       "BulkChange" + ch.EntityName + "Owner",
			Id:           "",
		},
	}
}

/**
 * Bulk Change owner of Record.
 */
func (ch *BulkChangeOwnerCommandHandler[M, C]) bulkChangeOwner(
	ctx context.Context,
	data map[string]account.MiniUser,
	hContext BulkChangeOwnerCommandHandlerContext[M, C]) (bool, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	for id, owner := range data {
		if record := ch.Repository.FindById(ctx, id); record != nil {
			hContext.SetOwner(record, owner)
			ch.Repository.Update(ctx, id, record)
		}
	}
	return true, nil
}
