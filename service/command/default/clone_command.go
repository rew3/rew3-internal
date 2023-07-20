package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/account/constants"
	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Clone command handler.
 * This handler can be used to clone record for any entity/model type.
 */
type CloneCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type CloneCommandHandlerContext[M common.Model, C command.Command] struct {
	SetOwner       func(*M, account.MiniUser)
	SetVisibility  func(*M, account.RecordVisibility)
	EmptyReference func(*M)
	EmptyFavorite  func(*M)
}

/**
 * Handle Command.
 */
func (ch *CloneCommandHandler[M, C]) Handle(ctx context.Context, idToClone string, hContext CloneCommandHandlerContext[M, C]) command.CommandResult {
	response, err := ch.clone(ctx, idToClone, hContext)
	return GenerateCmdResult[M](idToClone, response, err, "Clone"+ch.EntityName)
}

/**
 * Clone Record.
 */
func (ch *CloneCommandHandler[M, C]) clone(ctx context.Context, id string, hContext CloneCommandHandlerContext[M, C]) (*M, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		hContext.SetOwner(record, requestContext.User)
		hContext.EmptyReference(record)
		hContext.EmptyFavorite(record)
		visibility := account.RecordVisibility{}
		visibility.VisibilityType = constants.PRIVATE
		hContext.SetVisibility(record, visibility)
		return ch.Repository.Insert(ctx, record)
	} else {
		return nil, errors.New(ch.EntityName + " not found for given id")
	}
}
