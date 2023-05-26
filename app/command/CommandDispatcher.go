package command

import (
	"fmt"
	"reflect"
)

type CommandDispatcher struct {
	handlers map[string]CommandHandler
}

func NewCommandDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		handlers: make(map[string]CommandHandler),
	}
}

/**
 * Register Command handler. 
 */
func (dispatcher *CommandDispatcher) RegisterHandler(commandType string, handler CommandHandler) {
	dispatcher.handlers[commandType] = handler
}

/**
 * Dispatch Command. 
 */
func (d *CommandDispatcher) DispatchCommand(command Command) {
	commandType := reflect.TypeOf(command).Elem().Name()
	handler, exists := d.handlers[commandType]
	if exists {
		handler.HandleCommand(command)
	} else {
		fmt.Printf("No handler registered for command type: %s\n", commandType)
	}
}
