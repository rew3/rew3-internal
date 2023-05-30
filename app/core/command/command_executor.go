package command

import (
	"fmt"
	"reflect"

	s "github.com/rew3/rew3-base/app/core"
)

type CommandExecutor struct {
	serviceRegistry s.ServiceRegistry
}

func NewCommandExecutor(registry s.ServiceRegistry) *CommandExecutor {
	return &CommandExecutor{
		serviceRegistry: registry,
	}
}

/**
 * Execute Command.
 */
func (executor *CommandExecutor) Execute(command Command) {
	commandType := reflect.TypeOf(command).Elem().Name()
	handler, err := executor.serviceRegistry.GetCommandHandler(commandType)
	if err != nil {
		fmt.Printf("No handler registered for command type: %s\n", commandType)
	} else {
		handler.HandleCommand(command)
	}
}
