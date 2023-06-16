package executor

import (
	"context"
	"errors"
	"fmt"

	q "github.com/rew3/rew3-internal/service/query"
)

/**
 * Query Executor.
 */
type QueryExecutor struct {
	serviceRegistry *ServiceRegistry
}

func NewQueryExecutor(registry *ServiceRegistry) *QueryExecutor {
	return &QueryExecutor{
		serviceRegistry: registry,
	}
}

/**
 * Execute Query.
 */
func (executor *QueryExecutor) Execute(ctx context.Context, query q.Query) q.QueryResult {
	queryName := q.QueryName(query)
	controller, err := executor.serviceRegistry.GetQueryController(queryName)
	if err != nil {
		fmt.Printf("No handler/controller registered for query type: %s\n", queryName)
		return q.QueryResult{Error: errors.New("unable to process given query")}
	} else {
		resultChannel := q.NewQueryResultChannel()
		controller.Dispatch(ctx, query, resultChannel)
		result := <-resultChannel.Result
		return result
	}
}
