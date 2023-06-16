package command

import (
	"fmt"
	"time"

	s "github.com/rew3/rew3-internal/service/common/response"
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
		Result: make(chan CommandResult, 1),
	}
}
func (cs *CommandResultChannel) Send(data CommandResult) {
	select {
	case cs.Result <- data:
		fmt.Println("Data sent to Command Result Channel.")
	case <-time.After(time.Second):
		fmt.Println("Timeout reached while sending data to Command Result Channel.")
	}
	close(cs.Result)
}

/**
 * Parse the command result for given result type.
 * Provide default value in case of failure of parsing command result.
 */
func ParseCommandResult[T any](result CommandResult, defaultValue T) *s.ExecutionResult[T] {
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
		if data, ok := result.Response.Data.(T); ok {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         data,
			}
		} else if data, ok := result.Response.Data.(*T); ok {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         *data,
			}
		} else {
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
