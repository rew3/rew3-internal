package command

import "reflect"

/**
 * A super type for Command.
 * All command must use this root command using composition.
 */
type Command interface{}

func CommandName(command Command) string {
	return reflect.TypeOf(&command).Elem().Name()
}
