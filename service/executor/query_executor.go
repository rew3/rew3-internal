package executor

import (
	"context"
	"fmt"

	q "github.com/rew3/rew3-base/service/query"
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
	queryName := query.GetName()
	handler, err := executor.serviceRegistry.GetQueryHandler(queryName)
	if err != nil {
		fmt.Printf("No handler registered for query type: %s\n", queryName)
		return q.QueryResult{Error: "Unable to process given query."}
	} else {
		resultChannel := q.NewQueryResultChannel()
		handler.Handle(query, ctx, resultChannel)
		result := <-resultChannel.Result
		return result
	}
}
