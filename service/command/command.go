package command

/**
 * A super type for Command.
 * All command must use this root command using composition.
 */
type Command interface{
	GetName() string
}
