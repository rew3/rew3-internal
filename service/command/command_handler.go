package command

import (
	"context"

	response "github.com/rew3/rew3-internal/service/common/response"
)

type CommandHandler[T any] interface {
	Handle(ctx context.Context, cmd Command) response.ExecutionResult[T]
}
