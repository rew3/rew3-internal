package command

import "reflect"

/**
 * A super type for Command.
 * All command must use this root command using composition.
 */
type Command interface{}

/**
 * Get name of given command.
 * Note: you must provide command instance in pointer. e.g. ContactCommand. command := ContactCommand(nil); CommandName(&command)
 */
func CommandName(command Command) string {
	commandType := reflect.TypeOf(command)
	if commandType.Kind() == reflect.Ptr {
		commandType = commandType.Elem()
	}
	return commandType.Name()
}
