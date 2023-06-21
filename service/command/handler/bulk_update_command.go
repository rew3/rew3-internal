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
 * Bulk Update command handler.
 * This handler can be used to bulk update record for any entity/model type.
 */
type BulkUpdateCommandHandler[T common.Model] struct {
	entityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *BulkUpdateCommandHandler[T]) Handle(ctx context.Context,
	cmd command.Command,
	cmdToModels func(command.Command) (map[string]T, error),
	transformModel func(T) (T, error)) command.CommandResult {
	models, err := cmdToModels(cmd)
	if ok, cmdResult := HandleError(err, "BulkUpdate"+ch.entityName); !ok {
		return cmdResult
	}
	transformedModels := make(map[string]*T)
	for key, model := range models {
		transformed, err := transformModel(model)
		if err != nil {
			if ok, transformResult := HandleError(err, "BulkUpdate"+ch.entityName); !ok {
				return transformResult
			} else {
				transformedModels[key] = &transformed
			}
		}
	}
	_, errs := ch.bulkUpdate(ctx, transformedModels)
	if errs != nil {
		return command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.INTERNAL_SERVER_ERROR,
				Message:      err.Error(),
				Action:       "BulkUpdate" + ch.entityName,
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      ch.entityName + " successfully bulk updated",
			Action:       "BulkUpdate" + ch.entityName,
			Id:           "",
		},
	}
}

/**
 * Bulk Update Record.
 */
func (ch *BulkUpdateCommandHandler[T]) bulkUpdate(ctx context.Context, data map[string]*T) (bool, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	return ch.Repository.BulkUpdate(ctx, data)
}
