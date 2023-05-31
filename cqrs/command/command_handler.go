package command

type CommandHandler interface {
	HandleCommand(command Command)
}
