package context

import (
	"context"

	s "github.com/rew3/rew3-pkg/core/service/shared/request"
)

type contextKey int

const (
	requestContextKey contextKey = iota
)

// Store the RequestContext to context.
func WithRequestContext(ctx context.Context, reqContext s.RequestContext) context.Context {
	return context.WithValue(ctx, requestContextKey, reqContext)
}

// Retrieve the RequestContext from the context
func GetRequestContext(ctx context.Context) (s.RequestContext, bool) {
	reqContext, status := ctx.Value(requestContextKey).(s.RequestContext)
	return reqContext, status
}

// Retrieve the Context user id, return empty if not available.
func GetContextUserId(ctx context.Context) string {
	rc, hasValue := ctx.Value(requestContextKey).(s.RequestContext)
	if hasValue {
		return rc.User.Id
	} else {
		return ""
	}
}
