package core

import (
	c "github.com/rew3/rew3-base/app/core/command"
	q "github.com/rew3/rew3-base/app/core/query"
)

type ExecutorService struct {
	commandExecutor c.CommandExecutor
	queryExecutor   q.QueryExecutor
}

func NewExecutorService(registry *ServiceRegistry) *ExecutorService {
	return &ExecutorService{
		commandExecutor: c.NewCommandExecutor{registry},
		queryExecutor:   q.NewQueryExecutor{registry},
	}
}

func (service *ExecutorService) getCommandExecutor() *c.CommandExecutor {
	return &service.commandExecutor
}

func (service *ExecutorService) getQueryExecutor() *q.QueryExecutor {
	return &service.queryExecutor
}