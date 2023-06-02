package response

import (
	"github.com/rew3/rew3-base/app/common/constants"
)

type ExecutionResult[T any] struct {
	IsSuccessful bool
	Status       constants.StatusType
	Message      string
	Id           string
	Action       string
	Data         T
}
