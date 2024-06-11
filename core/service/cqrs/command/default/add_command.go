package handler

import (
	"context"
	"errors"

	"github.com/rew3/rew3-pkg/common/account"
	"github.com/rew3/rew3-pkg/common/account/constants"
	"github.com/rew3/rew3-pkg/core/service/cqrs/command"
	"github.com/rew3/rew3-pkg/core/service/shared"
	r "github.com/rew3/rew3-pkg/core/service/shared/response"
	c "github.com/rew3/rew3-pkg/core/service/shared/response/constants"
	"github.com/rew3/rew3-pkg/db/repository"

	rcUtil "github.com/rew3/rew3-pkg/core/service/utils/context"
)

/**
 * Add command handler.
 * This handler can be used to add record for any entity/model type.
 */
type AddCommandHandler[M shared.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type AddCommandHandlerContext[M shared.Model, C command.Command] struct {
	CmdToModel    func(*C) (M, error)
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
			Response: r.ErrorExecutionResult[M]("Add"+ch.EntityName, err.Error(), status),
		}
	}
	return command.CommandResult[M]{
		Response: r.SuccessExecutionResult[M]("Add"+ch.EntityName, "Successfully record added.", c.CREATED, *response),
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
	if hContext.GetOwner != nil && hContext.GetOwner(model).Id == "" {
		hContext.SetOwner(model, requestContext.User)
	}
	if hContext.GetVisibility != nil && hContext.GetVisibility(model).VisibilityType == "" {
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
