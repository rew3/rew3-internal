package query

import "context"

type QueryHandler interface {
	Handle(query Query, ctx context.Context)
}
