package query

import (
	"time"

	"github.com/rew3/rew3-internal/pkg/utils/logger"
	s "github.com/rew3/rew3-internal/service/common/response"
)

/**
 * Query Result.
 */
type QueryResult[T any] struct {
	Response s.ExecutionResult[T]
}

func (cr QueryResult[T]) Generify() QueryResult[interface{}] {
	return QueryResult[interface{}]{
		Response: cr.Response.Generify(),
	}
}

/**
 * Query result channel.
 */
type QueryResultChannel struct {
	Result chan QueryResult[interface{}]
}

func NewQueryResultChannel() *QueryResultChannel {
	return &QueryResultChannel{
		Result: make(chan QueryResult[interface{}], 1),
	}
}
func (cs *QueryResultChannel) Send(data QueryResult[interface{}]) {
	select {
	case cs.Result <- data:
	case <-time.After(time.Second):
		logger.Log().Error("Timeout reached while sending data to Query Result Channel.")
	}
	close(cs.Result)
}

/**
 * Parse the query result for given result type.
 * Provide default value in case of failure of parsing query result.
 */
func ParseQueryResult[T any](result *QueryResult[interface{}], defaultValue T) *s.ExecutionResult[T] {
	if !result.Response.IsSuccessful {
		return &s.ExecutionResult[T]{
			IsSuccessful: result.Response.IsSuccessful,
			Status:       result.Response.Status,
			Message:      result.Response.Message,
			Id:           result.Response.Id,
			Action:       result.Response.Action,
			Data:         defaultValue,
		}
	} else {
		if data, ok := result.Response.Data.(T); ok {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         data,
			}
		} else if data, ok := result.Response.Data.(*T); ok {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         *data,
			}
		} else {
			return &s.ExecutionResult[T]{
				IsSuccessful: result.Response.IsSuccessful,
				Status:       result.Response.Status,
				Message:      result.Response.Message,
				Id:           result.Response.Id,
				Action:       result.Response.Action,
				Data:         defaultValue,
			}
		}
	}
}
