package command

import (
	"context"
)

type CommandHandler[Response any, Cmd Command] interface {
	Handle(ctx context.Context, cmd Cmd) CommandResult[Response]
}
