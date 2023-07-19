package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/account/constants"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
)

/**
 * Add command handler.
 * This handler can be used to add record for any entity/model type.
 */
type AddCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type AddCommandHandlerContext[M common.Model, C command.Command] struct {
	CmdToModel    func(C) (M, error)
	GetOwner      func(*M) account.MiniUser
	SetOwner      func(*M, account.MiniUser)
	GetVisibility func(*M) account.RecordVisibility
	SetVisibility func(*M, account.RecordVisibility)
}

/**
 * Handle Command.
 */
func (ch *AddCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext AddCommandHandlerContext[M, C]) command.CommandResult {
	model, err := hContext.CmdToModel(cmd)
	if ok, cmdResult := HandleError(err, "Add"+ch.EntityName); !ok {
		return cmdResult
	}
	response, err := ch.add(ctx, &model, hContext)
	return GenerateCmdResult[M]("-", response, err, "Add"+ch.EntityName)
}

/**
 * Add Record.
 */
func (ch *AddCommandHandler[M, C]) add(ctx context.Context, model *M, hContext AddCommandHandlerContext[M, C]) (*M, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, errors.New("request context is not available")
	}
	if hContext.GetOwner(model).Id == "" {
		hContext.SetOwner(model, requestContext.User)
	}
	if hContext.GetVisibility(model).VisibilityType == "" {
		visibility := account.RecordVisibility{}
		visibility.VisibilityType = constants.PRIVATE
		hContext.SetVisibility(model, visibility)
	}
	return ch.Repository.Insert(ctx, model)
}
