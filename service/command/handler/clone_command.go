package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

	appCommon "github.com/rew3/rew3-internal/app/common"
	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Clone command handler.
 * This handler can be used to clone record for any entity/model type.
 */
type CloneCommandHandler[T common.Model] struct {
	entityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *CloneCommandHandler[T]) Handle(ctx context.Context, idToClone string) command.CommandResult {
	response, err := ch.clone(ctx, idToClone)
	return GenerateCmdResult[T](*response, err, "Clone"+ch.entityName)
}

/**
 * Clone Record.
 */
func (ch *CloneCommandHandler[T]) clone(ctx context.Context, id string) (*T, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		data := *record
		data.SetId("")
		data.SetMeta(appCommon.MetaInfo{})
		data.SetOwner(requestContext.User)
		data.SetReference(nil)
		return ch.Repository.Insert(ctx, &data)
	} else {
		return nil, errors.New(ch.entityName + " not found for given id")
	}
}
