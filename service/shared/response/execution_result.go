package response

import "github.com/rew3/rew3-internal/service/shared/response/constants"

type ExecutionResult[T any] struct {
	IsSuccessful bool
	Status       constants.StatusType
	Message      string
	Action       string
	Data         T
}

func (e *ExecutionResult[T]) Generify() ExecutionResult[interface{}] {
	return ExecutionResult[interface{}]{
		IsSuccessful: e.IsSuccessful,
		Status:       e.Status,
		Message:      e.Message,
		Action:       e.Action,
		Data:         e.Data,
	}
}

func ErrorExecutionResult[T any](action string, message string, status constants.StatusType) ExecutionResult[T] {
	return ExecutionResult[T]{
		IsSuccessful: false,
		Status:       status,
		Message:      message,
		Action:       action,
	}
}

func SuccessExecutionResult[T any](action string, message string, status constants.StatusType, data T) ExecutionResult[T] {
	return ExecutionResult[T]{
		IsSuccessful: true,
		Status:       status,
		Message:      message,
		Action:       action,
		Data:         data,
	}
}

// todo define ok, not_found, like scala. play. as ExecutionResult methods
func (c *ExecutionResult[T]) Ok(message string, data T) {

}
