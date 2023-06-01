package executor

type ExecutorService struct {
	commandExecutor *CommandExecutor
	queryExecutor   *QueryExecutor
}

func NewExecutorService(registry *ServiceRegistry) *ExecutorService {
	return &ExecutorService{
		commandExecutor: NewCommandExecutor(registry),
		queryExecutor:   NewQueryExecutor(registry),
	}
}

func (service *ExecutorService) CommandExecutor() *CommandExecutor {
	return service.commandExecutor
}

func (service *ExecutorService) QueryExecutor() *QueryExecutor {
	return service.queryExecutor
}
