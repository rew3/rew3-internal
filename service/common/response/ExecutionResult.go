package response

import (
	"github.com/rew3/rew3-internal/service/common/response/constants"
)

type ExecutionResult[T any] struct {
	IsSuccessful bool
	Status       constants.StatusType
	Message      string
	Id           string
	Action       string
	Data         T
}

func (e *ExecutionResult[T]) Generify() ExecutionResult[interface{}] {
	return ExecutionResult[interface{}]{
		IsSuccessful: e.IsSuccessful,
		Status:       e.Status,
		Message:      e.Message,
		Id:           e.Id,
		Action:       e.Action,
		Data:         e.Data,
	}
}

func ErrorExecutionResult[T any](id string, action string, message string, status constants.StatusType) ExecutionResult[T] {
	return ExecutionResult[T]{
		IsSuccessful: false,
		Status:       status,
		Message:      message,
		Id:           id,
		Action:       action,
	}
}

func SuccessExecutionResult[T any](id string, action string, message string, status constants.StatusType, data T) ExecutionResult[T] {
	return ExecutionResult[T]{
		IsSuccessful: true,
		Status:       status,
		Message:      message,
		Id:           id,
		Action:       action,
		Data:         data,
	}
}
