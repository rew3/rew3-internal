package internal

import (
	"context"
	"errors"

	"github.com/rew3/rew3-internal/db/repository/constants"
	mongoUtility "github.com/rew3/rew3-internal/db/utils"
	service "github.com/rew3/rew3-internal/service/shared/request"
	rcUtil "github.com/rew3/rew3-internal/service/utils/context"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * Request context resolver for Write operation.
 */
func HandleWrite[T any](ctx context.Context, operation func(service.RequestContext) (T, error)) (T, error) {
	if rc, isRcAvailable := rcUtil.GetRequestContext(ctx); isRcAvailable {
		return operation(rc)
	} else {
		var t T
		return t, errors.New("request context is not available")
	}
}

/**
 * Request context resolver for Read operation.
 */
func HandleRead[T any](ctx context.Context, operation func(service.RequestContext) T, defaultValue T) T {
	if rc, isRcAvailable := rcUtil.GetRequestContext(ctx); isRcAvailable {
		return operation(rc)
	} else {
		return defaultValue
	}
}

/**
 * Remove Internal fields form given document.
 */
func RemoveInternalFields(doc bson.D, skip ...constants.InternalField) bson.D {
	skipSet := make(map[string]bool)
	for _, field := range skip {
		skipSet[string(field)] = true
	}
	if _, skip := skipSet[string(constants.ID_FIELD)]; !skip {
		doc = mongoUtility.RemoveFieldFromBsonD(doc, string(constants.ID_FIELD))
	}
	if _, skip := skipSet[string(constants.META_FIELD)]; !skip {
		doc = mongoUtility.RemoveFieldFromBsonD(doc, string(constants.META_FIELD))
	}
	return doc
}
