package command

import (
	"fmt"
	"time"

	s "github.com/rew3/rew3-internal/service/common/response"
)

/**
 * Command Result.
 */
type CommandResult[T any] struct {
	Response s.ExecutionResult[T]
}

func (cr CommandResult[T]) Generify() CommandResult[interface{}] {
	return CommandResult[interface{}]{
		Response: cr.Response.Generify(),
	}
}

/**
 * Command result channel.
 */
type CommandResultChannel struct {
	Result chan CommandResult[interface{}]
}

func NewCommandResultChannel() *CommandResultChannel {
	return &CommandResultChannel{
		Result: make(chan CommandResult[interface{}], 1),
	}
}
func (cs *CommandResultChannel) Send(data CommandResult[interface{}]) {
	select {
	case cs.Result <- data:
		//fmt.Println("Data sent to Command Result Channel.")
	case <-time.After(time.Second):
		fmt.Println("Timeout reached while sending data to Command Result Channel.")
	}
	close(cs.Result)
}

/**
 * Parse the command result for given result type.
 * Provide default value in case of failure of parsing command result.
 */
func ParseCommandResult[T any](result *CommandResult[interface{}], defaultValue T) *s.ExecutionResult[T] {
	if !result.Response.IsSuccessful {
		return &s.ExecutionResult[T]{
			IsSuccessful: result.Response.IsSuccessful,
			Status:       result.Response.Status,
			Message:      result.Response.Message,
			Action:       result.Response.Action,
			Data:         defaultValue,
		}
	} else {
		if data, ok := result.Response.Data.(T); ok {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Action:       result.Response.Action,
				Data:         data,
			}
		} else if data, ok := result.Response.Data.(*T); ok {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Action:       result.Response.Action,
				Data:         *data,
			}
		} else {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Action:       result.Response.Action,
				Data:         defaultValue,
			}
		}
	}
}
