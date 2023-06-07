package executor

import (
	"fmt"

	c "github.com/rew3/rew3-internal/service/command"
	q "github.com/rew3/rew3-internal/service/query"
)

type ServiceRegistry struct {
	commandControllers map[string]c.CommandController
	queryControllers   map[string]q.QueryController
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		commandControllers: make(map[string]c.CommandController),
		queryControllers:   make(map[string]q.QueryController),
	}
}

/**
 * Register Command handler/controllers.
 */
func (r *ServiceRegistry) RegisterCommandController(command c.Command, controller c.CommandController) {
	commandName := c.CommandName(command)
	r.commandControllers[commandName] = controller
}

/**
 * Register Query handler/controllers.
 */
func (r *ServiceRegistry) RegisterQueryController(query q.Query, controller q.QueryController) {
	queryName := q.QueryName(query)
	r.queryControllers[queryName] = controller
}

/**
 * Get Command handler/controllers.
 */
func (r *ServiceRegistry) GetCommandController(command c.Command) (c.CommandController, error) {
	commandName := c.CommandName(command)
	handler, ok := r.commandControllers[commandName]
	if !ok {
		return nil, fmt.Errorf("command handler/controller not found for command: %s", commandName)
	}
	return handler, nil
}

/**
 * Get Query handler/controllers.
 */
func (r *ServiceRegistry) GetQueryController(query q.Query) (q.QueryController, error) {
	queryName := q.QueryName(query)
	handler, ok := r.queryControllers[queryName]
	if !ok {
		return nil, fmt.Errorf("query handler/controller not found for query: %s", queryName)
	}
	return handler, nil
}
