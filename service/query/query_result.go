package query

import "errors"

/**
 * Query Result.
 */
type QueryResult struct {
	Data  interface{}
	Error string
}

/**
 * Query result channel.
 */
type QueryResultChannel struct {
	Result chan QueryResult
}

func NewQueryResultChannel() QueryResultChannel {
	return QueryResultChannel{
		Result: make(chan QueryResult),
	}
}

/**
 * Parse the query result for given result type.
 * Provide default value in case of failure of parsing query result.
 */
func ParseQueryResult[T any](result *QueryResult, defaultValue T) (*T, error) {
	if result.Data == nil {
		return &defaultValue, errors.New(result.Error)
	} else {
		switch value := result.Data.(type) {
		case T:
			return &value, nil
		default:
			return &defaultValue, errors.New("Query result is invalid")
		}
	}
}
