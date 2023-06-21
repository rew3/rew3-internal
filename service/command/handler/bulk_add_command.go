package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/account/constants"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	bResponse "github.com/rew3/rew3-internal/service/common/response"
	bResponseConstant "github.com/rew3/rew3-internal/service/common/response/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"

	appCommon "github.com/rew3/rew3-internal/app/common"
	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Bulk Add command handler.
 * This handler can be used to bulk add record for any entity/model type.
 */
type BulkAddCommandHandler[T common.Model] struct {
	entityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *BulkAddCommandHandler[T]) Handle(ctx context.Context,
	cmd command.Command,
	cmdToModel func(command.Command) ([]T, error),
	transformModel func(T) (T, error)) command.CommandResult {
	models, err := cmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "BulkAdd"+ch.entityName); !ok {
		return cmdResult
	}
	var transformedModels []T
	for _, model := range models {
		transformed, err := transformModel(model)
		if err != nil {
			if ok, transformResult := HandleError(err, "BulkAdd"+ch.entityName); !ok {
				return transformResult
			} else {
				transformedModels = append(transformedModels, transformed)
			}
		}
	}
	_, errs := ch.bulkAdd(ctx, transformedModels)
	if errs != nil {
		return command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.INTERNAL_SERVER_ERROR,
				Message:      err.Error(),
				Action:       "BulkAdd" + ch.entityName,
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      ch.entityName + " successfully bulk added",
			Action:       "BulkAdd" + ch.entityName,
			Id:           "",
		},
	}
}

/**
 * Bulk Add Record.
 */
func (ch *BulkAddCommandHandler[T]) bulkAdd(ctx context.Context, data []T) (bool, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	var models []*T
	for _, model := range data {
		model.SetId(primitive.NewObjectID().Hex())
		if model.GetOwner().Id == "" {
			model.SetOwner(requestContext.User)
		}
		model.SetMeta(appCommon.MetaInfo{})
		if model.GetVisibility().VisibilityType == "" {
			visibility := account.RecordVisibility{}
			visibility.VisibilityType = constants.PRIVATE
			model.SetVisibility(visibility)
		}
		models = append(models, &model)
	}
	return ch.Repository.BulkInsert(ctx, models)
}
