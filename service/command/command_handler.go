package command

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, command Command, resultChannel CommandResultChannel)
}
