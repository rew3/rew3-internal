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
type AddCommandHandler[T common.Model] struct {
	EntityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *AddCommandHandler[T]) Handle(ctx context.Context,
	cmd command.Command,
	cmdToModel func(command.Command) (T, error),
	transformModel func(T) (T, error)) command.CommandResult {
	model, err := cmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "Add"+ch.EntityName); !ok {
		return cmdResult
	}
	transformedModel, err := transformModel(model)
	if ok, transformResult := HandleError(err, "Add"+ch.EntityName); !ok {
		return transformResult
	}
	response, err := ch.add(ctx, transformedModel)
	return GenerateCmdResult[T](*response, err, "Add"+ch.EntityName)
}

/**
 * Add Record.
 */
func (ch *AddCommandHandler[T]) add(ctx context.Context, data T) (*T, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	data.SetId(primitive.NewObjectID().Hex())
	if data.GetOwner().Id == "" {
		data.SetOwner(requestContext.User)
	}
	data.SetMeta(appCommon.MetaInfo{})
	if data.GetVisibility().VisibilityType == "" {
		visibility := account.RecordVisibility{}
		visibility.VisibilityType = constants.PRIVATE
		data.SetVisibility(visibility)
	}
	return ch.Repository.Insert(ctx, &data)
}
