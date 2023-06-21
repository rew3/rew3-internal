package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Change owner command handler.
 * This handler can be used to change owner of record for any entity/model type.
 */
type ChangeOwnerCommandHandler[T common.Model] struct {
	entityName string
	Repository repository.Repository[T]
}

/**
 * Handle Command.
 */
func (ch *ChangeOwnerCommandHandler[T]) Handle(ctx context.Context,
	cmd command.Command,
	idResolver func(command.Command) string,
	cmdToOwner func(command.Command) account.MiniUser) command.CommandResult {
	id := idResolver(cmd)
	owner := cmdToOwner(cmd)
	response, err := ch.changeOwner(ctx, id, owner)
	return GenerateCmdResult[T](*response, err, "Change"+ch.entityName+"Owner")
}

/**
 * Change owner of Record.
 */
func (ch *ChangeOwnerCommandHandler[T]) changeOwner(ctx context.Context, id string, owner account.MiniUser) (*T, error) {
	_, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if record := ch.Repository.FindById(ctx, id); record != nil {
		data := *record
		data.SetOwner(owner)
		return ch.Repository.Update(ctx, id, &data)
	} else {
		return nil, errors.New(ch.entityName + " not found for given id")
	}
}