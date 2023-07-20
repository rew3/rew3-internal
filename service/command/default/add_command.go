package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/app/account"
	"github.com/rew3/rew3-internal/app/account/constants"
	"github.com/rew3/rew3-internal/db/repository"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	r "github.com/rew3/rew3-internal/service/common/response"
	c "github.com/rew3/rew3-internal/service/common/response/constants"

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
	CmdToModel    func(*C) (M, error)
	GetId         func(*M) string
	GetOwner      func(*M) account.MiniUser
	SetOwner      func(*M, account.MiniUser)
	GetVisibility func(*M) account.RecordVisibility
	SetVisibility func(*M, account.RecordVisibility)
}

/**
 * Handle Command.
 */
func (ch *AddCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext AddCommandHandlerContext[M, C]) command.CommandResult[M] {
	model, err := hContext.CmdToModel(&cmd)
	if ok, cmdResult := HandleError[M](err, "Add"+ch.EntityName); !ok {
		return cmdResult
	}
	response, status, err := ch.add(ctx, &model, hContext)
	if err != nil {
		return command.CommandResult[M]{
			Response: r.ErrorExecutionResult[M]("-", "Add"+ch.EntityName, err.Error(), status),
		}
	}
	return command.CommandResult[M]{
		Response: r.SuccessExecutionResult[M](hContext.GetId(response), "Add"+ch.EntityName, "Successfully record added.", c.CREATED, *response),
	}
}

/**
 * Add Record.
 */
func (ch *AddCommandHandler[M, C]) add(ctx context.Context, model *M, hContext AddCommandHandlerContext[M, C]) (*M, c.StatusType, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return nil, c.FORBIDDEN, errors.New("request context is not available")
	}
	if hContext.GetOwner(model).Id == "" {
		hContext.SetOwner(model, requestContext.User)
	}
	if hContext.GetVisibility(model).VisibilityType == "" {
		visibility := account.RecordVisibility{}
		visibility.VisibilityType = constants.PRIVATE
		hContext.SetVisibility(model, visibility)
	}
	res, err := ch.Repository.Insert(ctx, model)
	if err != nil {
		return nil, c.INTERNAL_SERVER_ERROR, err
	}
	return res, c.CREATED, nil
}
