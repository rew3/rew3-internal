package command

import "context"

type CommandHandler interface {
	Handle(command Command, ctx context.Context)
}
