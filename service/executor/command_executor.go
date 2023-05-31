package executor

import (
	"fmt"

	c "github.com/rew3/rew3-base/service/command"
)

type CommandExecutor struct {
	serviceRegistry ServiceRegistry
}

func NewCommandExecutor(registry ServiceRegistry) *CommandExecutor {
	return &CommandExecutor{
		serviceRegistry: registry,
	}
}

/**
 * Execute Command.
 */
func (executor *CommandExecutor) Execute(command c.Command) {
	//commandType := reflect.TypeOf(command).Elem().Name()
	commandName := command.GetName()
	handler, err := executor.serviceRegistry.GetCommandHandler(commandName)
	if err != nil {
		fmt.Printf("No handler registered for command type: %s\n", commandName)
	} else {
		handler.Handle(command)
	}
}
