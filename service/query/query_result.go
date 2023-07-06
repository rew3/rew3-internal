package query

import (
	"errors"
	"fmt"
	"time"
)

/**
 * Query Result.
 */
type QueryResult struct {
	Data  interface{}
	Error error
}

/**
 * Query result channel.
 */
type QueryResultChannel struct {
	Result chan QueryResult
}

func NewQueryResultChannel() *QueryResultChannel {
	return &QueryResultChannel{
		Result: make(chan QueryResult, 1),
	}
}
func (cs *QueryResultChannel) Send(data QueryResult) {
	select {
	case cs.Result <- data:
	case <-time.After(time.Second):
		fmt.Println("Timeout reached while sending data to Query Result Channel.")
	}
	close(cs.Result)
}

/**
 * Parse the query result for given result type.
 * Provide default value in case of failure of parsing query result.
 */
func ParseQueryResult[T any](result *QueryResult, defaultValue T) (T, error) {
	if result.Data == nil {
		return defaultValue, result.Error
	} else {
		if data, ok := result.Data.(T); ok {
			return data, nil
		} else if data, ok := result.Data.(*T); ok {
			return *data, nil
		} else {
			return defaultValue, errors.New("Query result is invalid, unable to parse result")
		}
	}
}
