package shared

import (
	c "github.com/rew3/rew3-internal/service/shared/response/constants"
)

/**
 * Error type with status wrapped.
 */
type Error struct {
	Message string
	Status  c.StatusType
}

func NewError(message string, status c.StatusType) Error {
	return Error{
		Message: message,
		Status:  status,
	}
}
