package executor

import (
	"fmt"

	c "github.com/rew3/rew3-internal/service/command"
	q "github.com/rew3/rew3-internal/service/query"
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
func (r *ServiceRegistry) RegisterCommandHandler(command c.Command, handler c.CommandHandler) {
	commandName := c.CommandName(command)
	r.commandHandlers[commandName] = handler
}

/**
 * Register Query handler.
 */
func (r *ServiceRegistry) RegisterQueryHandler(query q.Query, handler q.QueryHandler) {
	queryName := q.QueryName(query)
	r.queryHandlers[queryName] = handler
}

/**
 * Get Command handler.
 */
func (r *ServiceRegistry) GetCommandHandler(command c.Command) (c.CommandHandler, error) {
	commandName := c.CommandName(command)
	handler, ok := r.commandHandlers[commandName]
	if !ok {
		return nil, fmt.Errorf("command handler not found for command: %s", commandName)
	}
	return handler, nil
}

/**
 * Get Query handler.
 */
func (r *ServiceRegistry) GetQueryHandler(query q.Query) (q.QueryHandler, error) {
	queryName := q.QueryName(query)
	handler, ok := r.queryHandlers[queryName]
	if !ok {
		return nil, fmt.Errorf("query handler not found for query: %s", queryName)
	}
	return handler, nil
}
