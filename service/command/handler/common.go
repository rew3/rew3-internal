package handler

import (
	"github.com/rew3/rew3-internal/pkg/utils/validator"
	"github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common"
	bResponse "github.com/rew3/rew3-internal/service/common/response"
	bResponseConstant "github.com/rew3/rew3-internal/service/common/response/constants"
)

func GenerateCmdResult[T common.Model](id string, data T, err error, action string) command.CommandResult {
	if err != nil {
		return command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.INTERNAL_SERVER_ERROR,
				Message:      err.Error(),
				Action:       action,
			},
		}
	}
	return command.CommandResult{
		Response: bResponse.ExecutionResult[interface{}]{
			IsSuccessful: true,
			Status:       bResponseConstant.CREATED,
			Message:      " successfully completed :" + action,
			Action:       action,
			Id:           id,
			Data:         data,
		},
	}
}

func ValidateModel[T common.Model](model T, action string) (bool, command.CommandResult) {
	dataValidator := validator.NewDataValidator()
	err := dataValidator.ValidatData(model)
	return HandleError(err, action)
}

func ValidateCommand(cmd interface{}, action string) (bool, command.CommandResult) {
	dataValidator := validator.NewDataValidator()
	err := dataValidator.ValidatData(cmd)
	return HandleError(err, action)
}

func HandleError(err error, action string) (bool, command.CommandResult) {
	if err != nil {
		return false, command.CommandResult{
			Response: bResponse.ExecutionResult[interface{}]{
				IsSuccessful: false,
				Status:       bResponseConstant.BAD_REQUEST,
				Message:      err.Error(),
				Action:       action,
			},
		}
	} else {
		return true, command.CommandResult{}
	}
}
