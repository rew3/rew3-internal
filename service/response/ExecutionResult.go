package response

import (
	"github.com/rew3/rew3-internal/app/common/constants"
)

type ExecutionResult[T any] struct {
	IsSuccessful bool
	Status       constants.StatusType
	Message      string
	Id           string
	Action       string
	Data         T
}

func ErrorExecutionResult[T any](id string, action string, message string, status constants.StatusType) *ExecutionResult[T] {
	return &ExecutionResult[T]{
		IsSuccessful: false,
		Status:       status,
		Message:      message,
		Id:           id,
		Action:       action,
	}
}

func SuccessExecutionResult[T any](id string, action string, message string, status constants.StatusType, data T) *ExecutionResult[T] {
	return &ExecutionResult[T]{
		IsSuccessful: true,
		Status:       status,
		Message:      message,
		Id:           id,
		Action:       action,
		Data:         data,
	}
}

func CreateExecutionResult[T any](id string, action string, message string, data T) *ExecutionResult[T] {
	return &ExecutionResult[T]{
		IsSuccessful: true,
		Status:       constants.CREATED,
		Message:      message,
		Id:           id,
		Action:       action,
		Data:         data,
	}
}

func UpdateExecutionResult[T any](id string, action string, message string, data T) *ExecutionResult[T] {
	return &ExecutionResult[T]{
		IsSuccessful: true,
		Status:       constants.OK,
		Message:      message,
		Id:           id,
		Action:       action,
		Data:         data,
	}
}

func DeleteExecutionResult[T any](id string, action string, message string) *ExecutionResult[T] {
	return &ExecutionResult[T]{
		IsSuccessful: true,
		Status:       constants.DELETED,
		Message:      message,
		Id:           id,
		Action:       action,
	}
}
