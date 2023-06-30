package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/account/constants"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	"go.mongodb.org/mongo-driver/bson/primitive"

	appCommon "github.com/rew3/rew3-internal/app/common"
	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Add command handler.
 * This handler can be used to add record for any entity/model type.
 */
type AddCommandHandler[W common.ModelWrapper, M common.Model, C command.Command] struct {
	EntityName      string
	Repository      repository.Repository[M]
	WrapperProvider func(*M) W
}

/**
 * Handle Command.
 */
func (ch *AddCommandHandler[W, M, C]) Handle(ctx context.Context,
	cmd C,
	cmdToModel func(C) (M, error),
	transformModel func(M) (M, error)) command.CommandResult {
	model, err := cmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "Add"+ch.EntityName); !ok {
		return cmdResult
	}
	transformedModel, err := transformModel(model)
	if ok, transformResult := HandleError(err, "Add"+ch.EntityName); !ok {
		return transformResult
	}
	data := &transformedModel
	wrapper := ch.WrapperProvider(data)
	response, err := ch.add(ctx, wrapper, &model)
	return GenerateCmdResult[M](wrapper.GetId(), response, err, "Add"+ch.EntityName)
}

/**
 * Add Record.
 */
func (ch *AddCommandHandler[W, M, C]) add(ctx context.Context, wrapper W, model *M) (*M, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
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
	return ch.Repository.Insert(ctx, model)
}
