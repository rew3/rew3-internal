package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/account/constants"
	appCommon "github.com/rew3/rew3-internal/app/common"
	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Clone command handler.
 * This handler can be used to clone record for any entity/model type.
 */
type CloneCommandHandler[W common.ModelWrapper, T common.Model, C command.Command] struct {
	EntityName      string
	Repository      repository.Repository[T]
	WrapperProvider func(*T) W
}

/**
 * Handle Command.
 */
func (ch *CloneCommandHandler[W, T, C]) Handle(ctx context.Context, idToClone string) command.CommandResult {
	response, err := ch.clone(ctx, idToClone)
	return GenerateCmdResult[T](idToClone, response, err, "Clone"+ch.EntityName)
}

/**
 * Clone Record.
 */
func (ch *CloneCommandHandler[W, T, C]) clone(ctx context.Context, id string) (*T, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		data := ch.WrapperProvider(record)
		data.SetId("")
		data.SetMeta(appCommon.MetaInfo{})
		data.SetOwner(requestContext.User)
		data.SetReference(nil)
		data.SetFavorite(false)
		visibility := account.RecordVisibility{}
		visibility.VisibilityType = constants.PRIVATE
		data.SetVisibility(visibility)
		return ch.Repository.Insert(ctx, record)
	} else {
		return nil, errors.New(ch.EntityName + " not found for given id")
	}
}
