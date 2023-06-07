package query

import "context"

type QueryController interface {
	Dispatch(ctx context.Context, query Query, resultChannel QueryResultChannel)
}
