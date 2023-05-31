package executor

import (
	"fmt"

	c "github.com/rew3/rew3-base/service/command"
	q "github.com/rew3/rew3-base/service/query"
)

type ServiceRegistry struct {
	commandHandlers map[string]c.CommandHandler
	queryHandlers   map[string]q.QueryHandler
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		commandHandlers: make(map[string]c.CommandHandler),
		queryHandlers:   make(map[string]q.QueryHandler),
	}
}

/**
 * Register Command handler.
 */
func (r *ServiceRegistry) RegisterCommandHandler(commandName string, handler c.CommandHandler) {
	r.commandHandlers[commandName] = handler
}

/**
 * Register Query handler.
 */
func (r *ServiceRegistry) RegisterQueryHandler(queryName string, handler q.QueryHandler) {
	r.queryHandlers[queryName] = handler
}

/**
 * Get Command handler.
 */
func (r *ServiceRegistry) GetCommandHandler(commandName string) (c.CommandHandler, error) {

	handler, ok := r.commandHandlers[commandName]
	if !ok {
		return nil, fmt.Errorf("command handler not found for command: %s", commandName)
	}
	return handler, nil
}

/**
 * Get Query handler.
 */
func (r *ServiceRegistry) GetQueryHandler(queryName string) (q.QueryHandler, error) {
	handler, ok := r.queryHandlers[queryName]
	if !ok {
		return nil, fmt.Errorf("query handler not found for query: %s", queryName)
	}
	return handler, nil
}
