package executor

import (
	"context"
	"time"

	"github.com/rew3/rew3-internal/pkg/utils/logger"
	c "github.com/rew3/rew3-internal/service/command"
	"github.com/rew3/rew3-internal/service/common/response"
	"github.com/rew3/rew3-internal/service/common/response/constants"
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
func (executor *CommandExecutor) Execute(ctx context.Context, command c.Command) c.CommandResult[interface{}] {
	commandName := c.CommandName(command)
	controller, err := executor.serviceRegistry.GetCommandController(commandName)
	if err != nil {
		message := "No handler/controller registered for command type: " + commandName
		logger.Log().Error(message)
		return c.CommandResult[interface{}]{Response: response.ErrorExecutionResult[interface{}]("-", commandName, message, constants.SERVICE_UNAVAILABLE)}
	} else {
		resultChannel := c.NewCommandResultChannel()
		controller.Dispatch(ctx, command, resultChannel)
		select {
		case result := <-resultChannel.Result:
			return result
		case <-time.After(30 * time.Second):
			message := "Timeout reached while receiving data by Command Executor."
			logger.Log().Error(message)
			return c.CommandResult[interface{}]{Response: response.ErrorExecutionResult[interface{}]("-", commandName, message, constants.SERVICE_UNAVAILABLE)}
		}
	}
}
