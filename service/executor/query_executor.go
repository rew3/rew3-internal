package executor

import (
	"context"
	"log"
	"time"

	"github.com/rew3/rew3-internal/service/common/response"
	"github.com/rew3/rew3-internal/service/common/response/constants"
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
func (executor *QueryExecutor) Execute(ctx context.Context, query q.Query) q.QueryResult[interface{}] {
	queryName := q.QueryName(query)
	controller, err := executor.serviceRegistry.GetQueryController(queryName)
	if err != nil {
		message := "No handler/controller registered for query type: " + queryName
		log.Println(message)
		return q.QueryResult[interface{}]{Response: response.ErrorExecutionResult[interface{}]("-", queryName, message, constants.SERVICE_UNAVAILABLE)}
	} else {
		resultChannel := q.NewQueryResultChannel()
		controller.Dispatch(ctx, query, resultChannel)
		select {
		case result := <-resultChannel.Result:
			return result
		case <-time.After(30 * time.Second):
			message := "Timeout reached while receiving data by Query Executor."
			log.Println(message)
			return q.QueryResult[interface{}]{Response: response.ErrorExecutionResult[interface{}]("-", queryName, message, constants.SERVICE_UNAVAILABLE)}
		}
	}
}
