package command

import (
	s "github.com/rew3/rew3-base/service/response"
)

type CommandResult[Entity any] struct {
	Success      bool
	Response     s.ExecutionResult[Entity]
	ErrorMessage string
}

func Success[Entity any](response s.ExecutionResult[Entity]) CommandResult[Entity] {
	return CommandResult[Entity]{
		Success:  true,
		Response: response,
	}
}

func Failure[Entity any](error string) CommandResult[Entity] {
	return CommandResult[Entity]{
		Success:      false,
		ErrorMessage: error,
	}
}

func (r CommandResult[Entity]) IsSuccess() bool {
	return r.Success
}

func (r CommandResult[Entity]) GetResponse() s.ExecutionResult[Entity] {
	return r.Response
}

func (r CommandResult[Entity]) GetErrorMessage() string {
	return r.ErrorMessage
}
