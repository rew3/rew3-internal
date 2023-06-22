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
type BulkChangeOwnerCommandHandler[T common.Model] struct {
	EntityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *BulkChangeOwnerCommandHandler[T]) Handle(ctx context.Context,
	cmd command.Command,
	cmdToOwner func(command.Command) map[string]account.MiniUser) command.CommandResult {
	owners := cmdToOwner(cmd)
	_, err := ch.bulkChangeOwner(ctx, owners)
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
func (ch *BulkChangeOwnerCommandHandler[T]) bulkChangeOwner(ctx context.Context, data map[string]account.MiniUser) (bool, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	for id, owner := range data {
		if record := ch.Repository.FindById(ctx, id); record != nil {
			data := *record
			data.SetOwner(owner)
			ch.Repository.Update(ctx, id, &data)
		}
	}
	return true, nil
}
