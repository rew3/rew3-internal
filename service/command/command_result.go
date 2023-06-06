package command

import (
	s "github.com/rew3/rew3-internal/service/response"
)

/**
 * Command Result.
 */
type CommandResult struct {
	Response s.ExecutionResult[interface{}]
}

/**
 * Command result channel.
 */
type CommandResultChannel struct {
	Result chan CommandResult
}

func NewCommandResultChannel() *CommandResultChannel {
	return &CommandResultChannel{
		Result: make(chan CommandResult),
	}
}

/**
 * Parse the command result for given result type.
 * Provide default value in case of failure of parsing command result.
 */
func ParseQueryResult[T any](result CommandResult, defaultValue T) *s.ExecutionResult[T] {
	if !result.Response.IsSuccessful {
		return &s.ExecutionResult[T]{
			IsSuccessful: result.Response.IsSuccessful,
			Status:       result.Response.Status,
			Message:      result.Response.Message,
			Id:           result.Response.Id,
			Action:       result.Response.Action,
			Data:         defaultValue,
		}
	} else {
		switch value := result.Response.Data.(type) {
		case T:
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         value,
			}
		default:
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         defaultValue,
			}
		}
	}
}
