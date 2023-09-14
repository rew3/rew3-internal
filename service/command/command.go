package command

import "reflect"

/**
 * A super type for Command.
 * All command must use this root command using composition.
 */
type Command interface{}

/**
 * Get name of given command.
 * Note: you must provide command instance in pointer if it is interface. e.g. ContactCommand. command := ContactCommand(nil); CommandName(&command)
 */
func CommandName(command Command) string {
	commandType := reflect.TypeOf(command)
	if commandType.Kind() == reflect.Ptr {
		commandType = commandType.Elem()
	}
	return commandType.Name()
}

/*
 * Get Parent of given concrete command (not interface). 
 * e.g. AddContact will return ContactCommand. 
 */
 func ParentCommandName(q Command) string {
	qType := reflect.TypeOf(q)
	if qType.Kind() == reflect.Struct {
		// Loop through the fields of the struct to find the embedded type
		for i := 0; i < qType.NumField(); i++ {
			field := qType.Field(i)
			if field.Anonymous {
				// Check if the embedded field implements Query interface
				if field.Type.Implements(reflect.TypeOf((*Command)(nil)).Elem()) {
					return field.Type.Name()
				}
			}
		}
	}
	return ""
}