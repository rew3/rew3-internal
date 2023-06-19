package repository

import (
	"context"
	"errors"

	rcUtil "github.com/rew3/rew3-internal/pkg/context"
	service "github.com/rew3/rew3-internal/service/common/request"
	mongoUtility "github.com/rew3/rew3-internal/db/utils"
	"github.com/rew3/rew3-internal/db/repository/constants"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * Request context resolver.
 */
func handle[T any](ctx context.Context, operation func(service.RequestContext) (T, error)) (T, error) {
	if rc, isRcAvailable := rcUtil.GetRequestContext(ctx); !isRcAvailable {
		return operation(rc)
	} else {
		var t T
		return t, errors.New("request context is not available")
	}
}

/**
 * Remove Internal fields form given document. 
 */
func removeInternalFields(doc bson.D) bson.D {
	doc = mongoUtility.RemoveFieldFromBsonD(doc, string(constants.ID_FIELD))
	doc = mongoUtility.RemoveFieldFromBsonD(doc, string(constants.META_FIELD))
	return doc
}