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
 * Bulk Add command handler.
 * This handler can be used to bulk add record for any entity/model type.
 */
type BulkAddCommandHandler[M common.Model, C command.Command] struct {
	EntityName string
	Repository repository.Repository[M]
}

type BulkAddCommandHandlerContext[M common.Model, C command.Command] struct {
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
			Response: r.ErrorExecutionResult[interface{}]("-", "BulkAdd"+ch.EntityName, errs.Error(), status),
		}
	}
	return command.CommandResult[interface{}]{
		Response: r.SuccessExecutionResult[interface{}]("-", "BulkAdd"+ch.EntityName, "Successfully bulk added records", c.OK, nil),
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
		if hContext.GetOwner(&model).Id == "" {
			hContext.SetOwner(&model, requestContext.User)
		}
		if hContext.GetVisibility(&model).VisibilityType == "" {
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
