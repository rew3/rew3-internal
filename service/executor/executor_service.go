package executor

type ExecutorService struct {
	commandExecutor CommandExecutor
	queryExecutor   QueryExecutor
}

func NewExecutorService(registry *ServiceRegistry) *ExecutorService {
	return &ExecutorService{
		commandExecutor: NewCommandExecutor{registry},
		queryExecutor:   NewQueryExecutor{registry},
	}
}

func (service *ExecutorService) getCommandExecutor() *CommandExecutor {
	return &service.commandExecutor
}

func (service *ExecutorService) getQueryExecutor() *QueryExecutor {
	return &service.queryExecutor
}
