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
func (ch *BulkDeleteCommandHandler[M, C]) Handle(ctx context.Context, idsToDelete []string) command.CommandResult[interface{}] {
	_, status, errs := ch.bulkDelete(ctx, idsToDelete)
	if errs != nil {
		return command.CommandResult[interface{}]{
			Response: r.ErrorExecutionResult[interface{}]("-", "BulkDelete"+ch.EntityName, errs.Error(), status),
		}
	}
	return command.CommandResult[interface{}]{
		Response: r.SuccessExecutionResult[interface{}]("-", "BulkDelete"+ch.EntityName, "Successfully bulk deleted records", c.OK, nil),
	}
}

/**
 * Bulk delete Record.
 */
func (ch *BulkDeleteCommandHandler[M, C]) bulkDelete(ctx context.Context, idsToDelete []string) (bool, c.StatusType, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, c.FORBIDDEN, errors.New("request context is not available")
	}
	_, err := ch.Repository.BulkDelete(ctx, idsToDelete)
	if err != nil {
		return false, c.INTERNAL_SERVER_ERROR, err
	}
	return true, c.OK, nil
}
