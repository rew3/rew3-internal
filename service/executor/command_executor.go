package executor

import (
	"context"
	"fmt"
	"time"

	c "github.com/rew3/rew3-internal/service/command"
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
	commandName := c.CommandName(command)
	controller, err := executor.serviceRegistry.GetCommandController(commandName)
	if err != nil {
		fmt.Printf("No handler/controller registered for command type: %s\n", commandName)
		return c.CommandResult{}
	} else {
		resultChannel := c.NewCommandResultChannel()
		controller.Dispatch(ctx, command, resultChannel)
		select {
		case result := <-resultChannel.Result:
			fmt.Println("Command result received by Command Executor.")
			return result
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout reached while receiving data by Command Executor.")
			return c.CommandResult{}
		}
	}
}
