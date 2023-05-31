package executor

import (
	"fmt"

	q "github.com/rew3/rew3-base/service/query"
)

type QueryExecutor struct {
	serviceRegistry ServiceRegistry
}

func NewQueryExecutor(registry ServiceRegistry) *QueryExecutor {
	return &QueryExecutor{registry}
}

/**
 * Execute Query.
 */
func (executor *QueryExecutor) Execute(query q.Query) {
	//queryType := reflect.TypeOf(query).Elem().Name()
	queryName := query.GetName()
	handler, err := executor.serviceRegistry.GetQueryHandler(queryName)
	if err != nil {
		fmt.Printf("No handler registered for query type: %s\n", queryName)
	} else {
		handler.Handle(query)
	}
}
