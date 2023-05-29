package service

import (
	"github.com/rew3/rew3-base/app/service/constants"
)

type ExecutionResult[T any] struct {
	IsSuccessful bool
	Status       constants.StatusType
	Message      string
	Id           string
	Action       string
	Data         T
}
