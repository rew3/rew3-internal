package handler

import (
	"github.com/rew3/rew3-pkg/core/service/cqrs/command"
	bResponse "github.com/rew3/rew3-pkg/core/service/shared/response"
	bResponseConstant "github.com/rew3/rew3-pkg/core/service/shared/response/constants"
)

func HandleError[T any](err error, action string) (bool, command.CommandResult[T]) {
	if err != nil {
		return false, command.CommandResult[T]{
			Response: bResponse.ExecutionResult[T]{
				IsSuccessful: false,
				Status:       bResponseConstant.BAD_REQUEST,
				Message:      err.Error(),
				Action:       action,
			},
		}
	} else {
		return true, command.CommandResult[T]{}
	}
}
