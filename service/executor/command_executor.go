package executor

import (
	"context"
	"fmt"

	c "github.com/rew3/rew3-base/service/command"
)

type CommandExecutor struct {
	serviceRegistry *ServiceRegistry
}

func NewCommandExecutor(registry *ServiceRegistry) *CommandExecutor {
	return &CommandExecutor{
		serviceRegistry: registry,
	}
}

/**
 * Execute Command.
 */
func (executor *CommandExecutor) Execute(ctx context.Context, command c.Command) c.CommandResult {
	commandName := command.GetName()
	handler, err := executor.serviceRegistry.GetCommandHandler(commandName)
	if err != nil {
		fmt.Printf("No handler registered for command type: %s\n", commandName)
		return c.CommandResult{}
	} else {
		resultChannel := c.NewCommandResultChannel()
		handler.Handle(command, ctx, *resultChannel)
		result := <-resultChannel.Result
		return result
	}
}
