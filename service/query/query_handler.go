package query

import "context"

type QueryHandler interface {
	Handle(ctx context.Context, query Query, resultChannel *QueryResultChannel)
}
