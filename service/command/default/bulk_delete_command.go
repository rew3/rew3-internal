package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	bResponse "github.com/rew3/rew3-internal/service/common/response"
	bResponseConstant "github.com/rew3/rew3-internal/service/common/response/constants"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Bulk Add command handler.
 * This handler can be used to bulk add record for any entity/model type.
 */
type BulkDeleteCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

/**
 * Handle Command.
 */
func (ch *BulkDeleteCommandHandler[M, C]) Handle(ctx context.Context, idsToDelete []string) command.CommandResult {
	_, errs := ch.bulkDelete(ctx, idsToDelete)
	if errs != nil {
		return command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.INTERNAL_SERVER_ERROR,
				Message:      errs.Error(),
				Action:       "BulkDelete" + ch.EntityName,
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      ch.EntityName + " successfully bulk deleted",
			Action:       "BulkDelete" + ch.EntityName,
			Id:           "",
		},
	}
}

/**
 * Bulk delete Record.
 */
func (ch *BulkDeleteCommandHandler[M, C]) bulkDelete(ctx context.Context, idsToDelete []string) (bool, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	return ch.Repository.BulkDelete(ctx, idsToDelete)
}
