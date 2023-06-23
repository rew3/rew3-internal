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
type BulkAddCommandHandler[W common.ModelWrapper, M common.Model, C command.Command] struct {
	EntityName      string
	Repository      repository.Repository[M]
	WrapperProvider func(*M) W
}

/**
 * Handle Command.
 */
func (ch *BulkAddCommandHandler[W, M, C]) Handle(ctx context.Context,
	cmd C,
	cmdToModel func(C) ([]M, error),
	transformModel func(M) (M, error)) command.CommandResult {
	models, err := cmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "BulkAdd"+ch.EntityName); !ok {
		return cmdResult
	}
	var transformedModels []M
	for _, model := range models {
		transformed, err := transformModel(model)
		if err != nil {
			if ok, transformResult := HandleError(err, "BulkAdd"+ch.EntityName); !ok {
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
				Action:       "BulkAdd" + ch.EntityName,
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      ch.EntityName + " successfully bulk added",
			Action:       "BulkAdd" + ch.EntityName,
			Id:           "",
		},
	}
}

/**
 * Bulk Add Record.
 */
func (ch *BulkAddCommandHandler[W, M, C]) bulkAdd(ctx context.Context, data []M) (bool, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, errors.New("request context is not available")
	}
	var models []*M
	for _, model := range data {
		pModel := &model
		wrapper := ch.WrapperProvider(pModel)
		wrapper.SetId(primitive.NewObjectID().Hex())
		if wrapper.GetOwner().Id == "" {
			wrapper.SetOwner(requestContext.User)
		}
		wrapper.SetMeta(appCommon.MetaInfo{})
		if wrapper.GetVisibility().VisibilityType == "" {
			visibility := account.RecordVisibility{}
			visibility.VisibilityType = constants.PRIVATE
			wrapper.SetVisibility(visibility)
		}
		models = append(models, pModel)
	}
	return ch.Repository.BulkInsert(ctx, models)
}
