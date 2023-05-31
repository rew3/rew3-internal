package query

import (
	"fmt"
	"reflect"

	s "github.com/rew3/rew3-base/cqrs"
)

type QueryExecutor struct {
	serviceRegistry s.ServiceRegistry
}

func NewQueryExecutor(registry s.ServiceRegistry) *QueryExecutor {
	return &QueryExecutor{registry}
}

/**
 * Execute Query.
 */
func (executor *QueryExecutor) Execute(query Query) {
	queryType := reflect.TypeOf(query).Elem().Name()
	handler, err := executor.serviceRegistry.GetQueryHandler(queryType)
	if err != nil {
		fmt.Printf("No handler registered for query type: %s\n", queryType)
	} else {
		handler.HandleQuery(query)
	}
}
