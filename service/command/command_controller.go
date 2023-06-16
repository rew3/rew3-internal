package command

import "context"

type CommandController interface {
	Dispatch(ctx context.Context, command Command, resultChannel *CommandResultChannel)
}
