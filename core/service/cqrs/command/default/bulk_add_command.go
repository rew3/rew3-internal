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
 * Bulk Add command handler.
 * This handler can be used to bulk add record for any entity/model type.
 */
type BulkAddCommandHandler[M shared.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type BulkAddCommandHandlerContext[M shared.Model, C command.Command] struct {
	CmdToModel    func(*C) ([]M, error)
	GetOwner      func(*M) account.MiniUser
	SetOwner      func(*M, account.MiniUser)
	GetVisibility func(*M) account.RecordVisibility
	SetVisibility func(*M, account.RecordVisibility)
}

/**
 * Handle Command.
 */
func (ch *BulkAddCommandHandler[M, C]) Handle(ctx context.Context, cmd C, hContext BulkAddCommandHandlerContext[M, C]) command.CommandResult[interface{}] {
	models, err := hContext.CmdToModel(&cmd)
	if ok, cmdResult := HandleError[interface{}](err, "BulkAdd"+ch.EntityName); !ok {
		return cmdResult
	}
	_, status, errs := ch.bulkAdd(ctx, models, hContext)
	if errs != nil {
		return command.CommandResult[interface{}]{
			Response: r.ErrorExecutionResult[interface{}]("BulkAdd"+ch.EntityName, errs.Error(), status),
		}
	}
	return command.CommandResult[interface{}]{
		Response: r.SuccessExecutionResult[interface{}]("BulkAdd"+ch.EntityName, "Successfully bulk added records", c.OK, nil),
	}
}

/**
 * Bulk Add Record.
 */
func (ch *BulkAddCommandHandler[M, C]) bulkAdd(ctx context.Context, data []M, hContext BulkAddCommandHandlerContext[M, C]) (bool, c.StatusType, error) {
	requestContext, isEcAvailable := rcUtil.GetRequestContext(ctx)
	if !isEcAvailable {
		return false, c.FORBIDDEN, errors.New("request context is not available")
	}
	var models []*M
	for _, model := range data {
		if hContext.GetOwner != nil && hContext.GetOwner(&model).Id == "" {
			hContext.SetOwner(&model, requestContext.User)
		}
		if hContext.GetVisibility != nil && hContext.GetVisibility(&model).VisibilityType == "" {
			visibility := account.RecordVisibility{}
			visibility.VisibilityType = constants.PRIVATE
			hContext.SetVisibility(&model, visibility)
		}
		models = append(models, &model)
	}
	_, err := ch.Repository.BulkInsert(ctx, models)
	if err != nil {
		return false, c.INTERNAL_SERVER_ERROR, err
	}
	return true, c.OK, nil
}
