package query

import (
	"context"
)

type QueryHandler[Response any, Q Query] interface {
	Handle(ctx context.Context, query Q) QueryResult[Response]
}
